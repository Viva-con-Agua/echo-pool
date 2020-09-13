package api

import (
	"log"
	"strings"
)

type (
	AccessList struct {
		Access map[string]map[string][]string `json:"access"`
	}
	// For user model
	Access struct {
		AccessUuid string `json:"access_uuid,omitempty"`
		AccessName string `json:"name" validate:"required"`
		ModelUuid  string `json:"model_uuid,omitempty"`
		ModelName  string `json:"model_name,omitempty"`
		ModelType  string `json:"model_type,omitempty"`
		Created    int64  `json:"created" validate:"required"`
	}
	//AccessList map[string][]Access

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
	UserSession struct {
		Uuid          string                         `json:"uuid"`
		Email         string                         `json:"email"`
		Confirmed     int                            `json:"confirmed"`
		Access        map[string]map[string][]string `json:"access"`
		PrivacyPolicy int                            `json:"privacy_policy"`
		Updated       int64                          `json:"updated"`
		Created       int64                          `json:"created"`
		Additional    Additional                     `json:"additional"`
	}

	AccessToken struct {
		AccessToken string `json:"access_token"`
	}
	Profile struct {
		Uuid      string    `json:"uuid" `
		Avatar    Avatar    `json:"avatar" `
		FirstName string    `json:"first_name" `
		LastName  string    `json:"last_name" `
		FullName  string    `json:"full_name" `
		Mobile    string    `json:"mobile_phone"`
		Birthdate int       `json:"birthdate"`
		Gender    string    `json:"gender" `
		Addresses []Address `json:"addresses"`
		Updated   int       `json:"updated" `
		Created   int       `json:"created" `
	}
	Avatar struct {
		Url  string `json:"url"`
		Data string `json:"data"`
		Type string `json:"type"`
	}
	Address struct {
		Uuid       string `json:"uuid" `
		Primary    int    `json:"primary" `
		Street     string `json:"street" `
		Additional string `json:"additional"`
		Zip        string `json:"zip" `
		City       string `json:"city" `
		Country    string `json:"country" `
		GoogleId   string `json:"google_id" `
		Updated    int    `json:"updated" `
		Created    int    `json:"created"`
	}
	AddressList []Address
	M           map[string]interface{}

	UserRequest struct {
		Uuid map[string]Additional `json:"uuid" validate:"required"`
	}
	Additional map[string]interface{}
)

func (u *User) Session() *UserSession {
	us := new(UserSession)
	us.Uuid = u.Uuid
	us.Email = u.Email
	us.Access = u.Access.Access
	us.Confirmed = u.Confirmed
	us.Updated = u.Updated
	us.Created = u.Created
	return us
}

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
