package api

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

func (l PermissionList) Contains(r Access) bool {
	for _, v := range l {
		if v.Name == r.AccessName {
			return true
		}
	}
	return false
}
