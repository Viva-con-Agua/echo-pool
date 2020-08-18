package auth

type (
	Access struct {
		Uuid        string `json:"uuid" validate:"required"`
		AccessName  string `json:"access_name" validate:"required"`
		ServiceName string `json:"service_name" validate:"required"`
		ModelUuid   string `json:"model_uuid"`
		ModeName    string `json:"model_name"`
		ModelType   string `json:"model_type"`
		Created     int    `json:"created" validate:"required"`
	}

	User struct {
		Uuid      string   `json:"uuid"`
		Email     string   `json:"email"`
		Name      string   `json:"name"`
		Confirmed int      `json:"confirmed"`
		Access    []Access `json:"access"`
		Profile   Profile  `json:"profile"`
		Updated   int      `json:"updated"`
		Created   int      `json:"created"`
	}
	AccessToken struct {
		AccessToken string `json:"access_token"`
	}
	Profile struct {
		Uuid      string    `json:"uuid" validate:"required"`
		Avatar    Avatar    `json:"avatar" validate:"required"`
		FirstName string    `json:"first_name" validate:"required"`
		LastName  string    `json:"last_name" validate:"required"`
		FullName  string    `json:"full_name" validate:"required"`
		Mobile    string    `json:"mobile_phone" validate:"required"`
		Birthdate int       `json:"birthdate" validate:"required"`
		Gender    string    `json:"gender" validate:"required"`
		Addresses []Address `json:"addresses" validate:"required"`
		Updated   int       `json:"updated" validate:"required"`
		Created   int       `json:"created" validate:"required"`
	}
	Avatar struct {
		Url  string `json:"url" validate:"required"`
		Data string `json:"data" validate:"required"`
		Type string `json:"type" validate:"required"`
	}
	Address struct {
		Uuid       string `json:"uuid" validate:"required"`
		Primary    int    `json:"primary" validate:"required"`
		Street     string `json:"street" validate:"required"`
		Additional string `json:"additional" validate:"required"`
		Zip        string `json:"zip" validate:"required"`
		City       string `json:"city" validate:"required"`
		Country    string `json:"country" validate:"required"`
		GoogleId   string `json:"google_id" validate:"required"`
		Updated    int    `json:"updated" validate:"required"`
		Created    int    `json:"created" validate:"required"`
	}
	M map[string]interface{}
)
