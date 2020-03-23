package pool

import (
  "encoding/gob"
  "bytes"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/rbcervilla/redisstore"
)

type (
	PoolUser struct {
		Uuid string `json:"uuid"`
		Email string `json:"email"`
		Name string `json:"name"`
		Roles []PoolRole `json:"roles"`
	}
	M map[string]interface{}
)

func RedisSession() echo.MiddlewareFunc {
	client := redis.NewClient(&redis.Options{
		Addr: Config.Redis.Url,
	})

	redis, err := redisstore.NewRedisStore(client)

	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}
  gob.Register(&PoolUser{})
  gob.Register(&M{})
	log.Println("Redis successfully connected!")
	return session.Middleware(redis)
}

func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["valid"] != nil {
		
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, Unauthorized())
	}
}

func GetAccessToken(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil

}

func DropsSessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["valid"] != nil {
			val := sess.Values["user"]
			var user = &PoolUser{}
			user, _ = val.(*PoolUser)
			if in := LogoutList.Contains(user.Uuid); !in {
				return next(c)
			}
			LogoutList = LogoutList.Delete(user.Uuid)
			return echo.NewHTTPError(http.StatusUnauthorized, Unauthorized())
		}
		return echo.NewHTTPError(http.StatusUnauthorized, Unauthorized())
	}
}

func SetSession(c echo.Context, user *PoolUser, token *AccessToken) {
	LogoutList.Delete(user.Uuid)
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	if sess.Values["valid"] == nil {
		sess.Values["valid"] = true
		sess.Values["token"] = token.AccessToken
		sess.Values["user"] = &user
		sess.Save(c.Request(), c.Response())
	}
}

func DeleteSession(c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Values["valid"] = nil
	sess.Save(c.Request(), c.Response())
}



