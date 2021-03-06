package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Viva-con-Agua/echo-pool/config"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SetSession(c echo.Context, user *UserSession) {
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
	sessionUser, _ := json.Marshal(user)
	sess.Values["valid"] = true
	sess.Values["user"] = &sessionUser
	sess.Save(c.Request(), c.Response())
}

func GetUser(c echo.Context) (u *UserSession, contains bool) {
	sess, _ := session.Get("session", c)
	val := sess.Values["user"]
	var user []byte
	user, contains = val.([]byte)
	if contains == false {
		return nil, contains
	}
	json.Unmarshal(user, &u)
	return u, true

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
