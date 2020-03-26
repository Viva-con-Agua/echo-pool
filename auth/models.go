package auth 

type (
	Role struct {
		Name string `json:"name"`
		CrewId string `json:"crewId"`
		CrewName string `json:"crewName"`
		Pillar string `json:"pillar"`
	}
	User struct {
		Uuid string `json:"uuid"`
		Email string `json:"email"`
		Name string `json:"name"`
		Roles []Role `json:"roles"`
	}
	AccessToken struct {
		AccessToken string `json:"access_token"`
	}
	M map[string]interface{}
)
