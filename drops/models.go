package drops

import "github.com/Viva-con-Agua/echo-pool/auth"

type (
	City struct {
		Country string `json:"country"`
		Name string `json:"name"` 
	}
	Crew struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Cities []City `json:"cities"`
	}
	Pillar struct {
		Pillar string `json:"pillar"`
	}
	Role struct {
		Crew Crew `json:"crew"`
		Name string `json:"name"`
		Pillar Pillar `json:"pillar"`
	}
	Role2 struct {
		Role string `json:"role"`
	}
	Supporter struct {
		FullName string `json:"fullName"`
		Roles []Role `json:"roles"`
	}
	Profile struct {
		Email string `json:"email"`
		Supporter Supporter `json:"supporter"`
	}
	Address struct {
		Additional string `json:"additional"`
		City string `json:"city"`
		Country string `json:"country"`
		PublicId string `json:"publicId"`
		Street string `json:"street"`
		Zip string `json:"zip"`
	}
	DropsUser struct {
		Id	string `json:"id"`
		Profiles []Profile `json:"profiles"`
		Roles	[]Role2 `json:"roles"`
	}


)

func (u* DropsUser) PoolUser() *auth.User {
	user := new(auth.User)
	user.Uuid = u.Id
	for _, p := range u.Profiles {
		user.Email = p.Email
		user.Name = p.Supporter.FullName
		var roles []auth.Role
		for _, r := range u.Roles {
			role := new(auth.Role)
			role.Name = r.Role
			roles = append(roles, *role)
		}
		var ddRoles = p.Supporter.Roles
		for _, r := range ddRoles {
			role := new(auth.Role)
			role.Name = r.Name
			role.Pillar = r.Pillar.Pillar
			role.CrewId = r.Crew.Id
			role.CrewName = r.Crew.Name
			roles = append(roles, *role)
		}
		user.Roles = roles
	}
	return user
} 


