package auth

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

func (l PermissionList) Contains(r Role) bool {
	for _, v := range l {
		if v.Name == r.Name {
			if v.Pillar == r.Pillar {
				return true
			}
		}
	}
	return false
}




