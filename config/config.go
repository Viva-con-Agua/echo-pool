package config

import (
	"github.com/jinzhu/configor"
)

var (
	Config = struct {
		Drops struct {
			Url struct {
				Code     string
				Redirect string
				Token    string
				User     string
			}
			Oauth struct {
				ClientId string
			}
		}
		Redis struct {
			Url string
		}
		Nats struct {
			Url string
		}
	}{}
)

func LoadConfig() {
	configor.Load(&Config, "config/config.yml")
}
