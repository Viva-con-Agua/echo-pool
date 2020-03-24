package auth 

type (
	PoolRole struct {
		Name string `json:"name"`
		CrewId string `json:"crewId"`
		CrewName string `json:"crewName"`
		Pillar string `json:"pillar"`
	}
	PoolUser struct {
		Uuid string `json:"uuid"`
		Email string `json:"email"`
		Name string `json:"name"`
		Roles []PoolRole `json:"roles"`
	}
	AccessToken struct {
		AccessToken string `json:"access_token"`
	}
	M map[string]interface{}
)
