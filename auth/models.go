package auth

import (
	"log"
	"strings"
)

type (

	// For user model
	Access struct {
		AccessUuid string `json:"access_uuid,omitempty"`
		AccessName string `json:"name" validate:"required"`
		ModelUuid  string `json:"model_uuid,omitempty"`
		ModelName  string `json:"model_name,omitempty"`
		ModelType  string `json:"model_type,omitempty"`
		Created    int64  `json:"created" validate:"required"`
	}
	AccessList map[string][]Access

	User struct {
		Uuid       string     `json:"uuid"`
		Email      string     `json:"email"`
		Confirmed  int        `json:"confirmed"`
		Access     AccessList `json:"access"`
		Profile    Profile    `json:"profile"`
		Updated    int64      `json:"updated"`
		Created    int64      `json:"created"`
		Additional Additional `json:"additional"`
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
	AddressList []Address
	M           map[string]interface{}

	UserRequest struct {
		Uuid map[string]Additional `json:"uuid" validate:"required"`
	}
	Additional map[string]interface{}
)

func (f *UserRequest) Filter() string {
	if f != nil {
		filter := "WHERE "
		for key, _ := range f.Uuid {
			log.Print(key)
			filter = filter + "u.uuid = '" + key + "' OR "
		}
		filter = strings.TrimSuffix(filter, "OR ")
		return filter
	} else {
		return ""
	}
}

func (req *UserRequest) Additional(u []User) []User {
	var list []User
	for _, val := range u {
		log.Print(req)
		val.Additional = (req.Uuid)[val.Uuid]
		list = append(list, val)
	}
	return list

}

func (list *AddressList) Distinct() *AddressList {
	r := make(AddressList, 0, len(*list))
	m := make(map[Address]bool)
	for _, val := range *list {
		if _, ok := m[val]; !ok {
			m[val] = true
			r = append(r, val)
		}
	}
	return &r
}
