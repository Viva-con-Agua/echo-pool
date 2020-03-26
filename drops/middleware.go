package drops

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/Viva-con-Agua/echo-pool/resp"
	"github.com/Viva-con-Agua/echo-pool/auth"
)

/**
	* Middleware for drops oauth2 handshake.
	* Check whether the user has a session and his UUID is 
	* not in the LogoutList.
	* Drops handle logout via Nats message broker
	*/

func DropsSessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		// check if session is valid
		if sess.Values["valid"] != nil {
			// get user from session storage
			user, contains := auth.GetUser(c)
			// if there is no user, return Unauthorized
			if contains == false {
				auth.DeleteSession(c)
				return echo.NewHTTPError(http.StatusUnauthorized, resp.Unauthorized())
			}
			// if user not in LogoutList, go to next controller/middleware
			if in := LogoutList.Contains(user.Uuid); !in {
				return next(c)
			}
			//if user is Unauthorized, delete the UUID from List
			LogoutList = LogoutList.Delete(user.Uuid)
		}
		// delete session
		auth.DeleteSession(c)
		return echo.NewHTTPError(http.StatusUnauthorized, resp.Unauthorized())
	}
}
