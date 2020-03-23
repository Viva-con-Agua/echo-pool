package pool

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/labstack/echo-contrib/session"
)
	
type (
	Permission struct {
		Name string
		Pillar string
	}
	PermissionList []Permission
)

var (
	NoRole = &Permission{Name: "foo", Pillar: "bar"}
	Admin = &Permission{Name: "admin", Pillar: ""}
)
 
func (l PermissionList) Add(p Permission) PermissionList {
	return append(l, p)
}

func (l PermissionList) Contains(r PoolRole) bool {
	for _, v := range l {
		if v.Name == r.Name {
			if v.Pillar == r.Pillar {
				return true
			}
		}
	}
	return false
}



func CheckPermission(permission *PermissionList) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			val := sess.Values["user"]
			var user = &PoolUser{}
			user, _ = val.(*PoolUser)
			for _, v := range user.Roles {
				if permission.Contains(v) {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusUnauthorized, Unauthorized())
		}
	}
}
