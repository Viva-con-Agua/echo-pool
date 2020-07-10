package auth

type (
	AccessUser struct {
		Uuid        string `json:"uuid" validate:"required"`
		RoleUuid    string `json:"role_uuid" validate:"required"`
		RoleName    string `json:"role_name" validate:"required"`
		ServiceName string `json:"service_name" validate:"required"`
		ModelUuid   string `json:"model_uuid"`
		ModeName    string `json:"model_name"`
		Created     int    `json:"created" validate:"required"`
	}
	Role struct {
		Name     string `json:"name"`
		CrewId   string `json:"crewId"`
		CrewName string `json:"crewName"`
		Pillar   string `json:"pillar"`
	}
	User struct {
		Uuid      string       `json:"uuid"`
		Email     string       `json:"email"`
		Name      string       `json:"name"`
		Confirmed int          `json:"confirmed"`
		Access    []AccessUser `json:"access"`
		Updated   int          `json:"updated"`
		Created   int          `json:"created"`
	}
	AccessToken struct {
		AccessToken string `json:"access_token"`
	}
	M map[string]interface{}
)
