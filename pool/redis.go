package pool

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/rbcervilla/redisstore"
)

func RedisSession(address string) echo.MiddlewareFunc {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	redis, err := redisstore.NewRedisStore(client)

	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}
	log.Println("Redis successfully connected!")
	return session.Middleware(redis)
}

func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["valid"] == nil || sess.Values["valid"] == false {
			return echo.NewHTTPError(http.StatusUnauthorized, Unauthorized())
		}
		return next(c)
	}
}
