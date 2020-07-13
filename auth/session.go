package auth

import (
	"net/http"

	"github.com/Viva-con-Agua/echo-pool/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func SetSession(c echo.Context, user *User, token *AccessToken) {
	secure := true
	if config.Config.Cookie.Secure == "false" {
		secure = false
	}
	sameSite := http.SameSiteNoneMode
	if config.Config.Cookie.Samesite == "lax" {
		sameSite = http.SameSiteLaxMode
	}
	if config.Config.Cookie.Samesite == "none" {
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
	if sess.Values["valid"] == nil {
		sess.Values["valid"] = true
		sess.Values["token"] = token.AccessToken
		sess.Values["user"] = &user
		sess.Save(c.Request(), c.Response())
	}
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
	sess.Values["token"] = nil
	sess.Values["user"] = nil
	sess.Save(c.Request(), c.Response())
}
