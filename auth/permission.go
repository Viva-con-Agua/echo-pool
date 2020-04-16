package auth

type (
	Permission struct {
		Name string
	}
	PermissionList []Permission
)

var (
	NoRole = &Permission{Name: "foo"}
	Admin  = &Permission{Name: "admin"}
)

func (l PermissionList) Add(p Permission) PermissionList {
	return append(l, p)
}

func (l PermissionList) Contains(r AccessUser) bool {
	for _, v := range l {
		if v.Name == r.RoleName {
			return true
		}
	}
	return false
}
