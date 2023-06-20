package config

import (
	"github.com/jinzhu/configor"
)

var Conf = struct {
	Debug bool
	// actual type is zerolog.Level
	LogLevel int `yaml:"log-level"`
	Bind     string
	Bahn     struct {
		ApiKey  string `yaml:"api-key"`
		BaseUrl string `yaml:"base-url"`
	}
	Db struct {
		Host     string
		User     string
		Password string
		DbName   string `yaml:"db-name"`
	}
	Jwt struct {
		Cookie         string
		ExpirationDays int `yaml:"expiration-days"`
		Secret         string
	}
	Requests struct {
		TimeoutSeconds int      `yaml:"timeout-seconds"`
		CorsOrigins    []string `yaml:"cors-origins"`
	}
	PushNotifications struct {
		VapidKeys struct {
			Public  string
			Private string
		} `yaml:"vapid-keys"`
		Subject string
		Ttl     int
	} `yaml:"push-notifications"`
}{}

func init() {
	if err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&Conf, "config.yml"); err != nil {
		panic(err)
	}

	if Conf.PushNotifications.VapidKeys.Public == "" {
		panic("Public vapid key not set. Did you forgot to specify a configor env using 'CONFIGOR_ENV=dev'?")
	}
}
