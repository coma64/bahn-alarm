package config

import (
	"github.com/jinzhu/configor"
)

var Conf = struct {
	Debug bool
	Bind  string
	Bahn  struct {
		ApiKey  string `yaml:"api-key"`
		BaseUrl string `yaml:"base-url"`
	}
	Db struct {
		Dsn string
	}
	Jwt struct {
		Cookie         string
		ExpirationDays int `yaml:"expiration-days"`
		Secret         string
	}
	RequestTimeoutSeconds int `yaml:"request-timeout-seconds"`
	PushNotifications     struct {
		VapidKeys struct {
			Public  string
			Private string
		} `yaml:"vapid-keys"`
		Subject string
	} `yaml:"push-notifications"`
}{}

func init() {
	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&Conf, "config.yml"); err != nil {
		panic(err)
	}
}
