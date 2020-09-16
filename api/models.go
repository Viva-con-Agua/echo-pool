package api

import (
	"log"
	"strings"
)

type (
	UserSession struct {
		Uuid          string                         `json:"uuid"`
		Email         string                         `json:"email"`
		Confirmed     bool                           `json:"confirmed"`
		PrivacyPolicy bool                           `json:"privacy_policy"`
		Country       string                         `json:"country"`
		Access        map[string]map[string][]string `json:"access"`
		Additional    Additional                     `json:"additional"`
		Updated       int64                          `json:"updated"`
		Created       int64                          `json:"created"`
	}
	AccessList struct {
		Access map[string]map[string][]string `json:"access"`
	}

	// Model can be every model
	Model struct {
		Uuid    string `json:"uuid" validate:"required"`
		Name    string `json:"name" validate:"required"`
		Service string `json:"service" validate:"required"`
		Creator string `json:"owner" validate:"required"`
		Created int64  `json:"created" validate:"required"`
	}
	ModelCreate struct {
		Uuid    string `json:"uuid" validate:"required"`
		Name    string `json:"name" validate:"required"`
		Service string `json:"service" validate:"required"`
		Creator string `json:"owner" validate:"required"`
	}

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

func (m_create *ModelCreate) Model(created int64) *Model {
	m := new(Model)
	m.Uuid = m_create.Uuid
	m.Name = m_create.Name
	m.Service = m_create.Service
	m.Creator = m_create.Creator
	m.Created = created
	return m
}
