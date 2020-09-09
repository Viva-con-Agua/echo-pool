package auth

import (
	"net/http"
	"os"

	"github.com/Viva-con-Agua/echo-pool/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func SetSession(c echo.Context, user *User) {
	secure := true
	if os.Getenv("COOKIE_SECURE") == "false" {
		secure = false
	}
	sameSite := http.SameSiteNoneMode
	if os.Getenv("SAME_SITE") == "lax" {
		sameSite = http.SameSiteLaxMode
	}
	if os.Getenv("SAME_SITE") == "none" {
		sameSite = http.SameSiteNoneMode
	}
	if config.Config.Cookie.Samesite == "strict" {
		sameSite = http.SameSiteStrictMode
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: sameSite,
		Secure:   secure,
	}
	sess.Values["valid"] = true
	sess.Values["user"] = &user
	sess.Save(c.Request(), c.Response())
}

func GetUser(c echo.Context) (u *User, contains bool) {
	sess, _ := session.Get("session", c)
	val := sess.Values["user"]
	var user = &User{}
	user, contains = val.(*User)
	if contains == false {
		return nil, contains
	}
	return user, true

}

func DeleteSession(c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Values["valid"] = nil
	sess.Values["user"] = nil
	sess.Save(c.Request(), c.Response())
}
