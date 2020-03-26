package auth

import (
	"log"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)






func SetSession(c echo.Context, user *PoolUser, token *AccessToken) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	log.Print(sess.Values["valid"])
	if sess.Values["valid"] == nil {
		sess.Values["valid"] = true
		sess.Values["token"] = token.AccessToken
		sess.Values["user"] = &user
		log.Print(user)
		log.Print(sess.Values["user"])
		sess.Save(c.Request(), c.Response())
	}
}

func GetUser(c echo.Context) (u *PoolUser, contains bool){
	sess, _ := session.Get("session", c)
	val := sess.Values["user"]
	var user = &PoolUser{}
	user, contains = val.(*PoolUser)
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


