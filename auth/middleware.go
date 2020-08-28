package auth

import (
	"bytes"
	"encoding/gob"
	"net/http"

	"github.com/Viva-con-Agua/echo-pool/resp"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if sess.Values["valid"] != nil {

			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, resp.Unauthorized())
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

/*
func CheckPermission(permission *PermissionList) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			val := sess.Values["user"]
			var user = &User{}
			user, _ = val.(*User)
			for _, v := range user.Access {
				if permission.Contains(v) {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusUnauthorized, resp.Unauthorized())
		}
	}
}*/
