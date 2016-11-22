package config

import (
	"github.com/vrischmann/envconfig"
	"log"
)

var conf *Config

/*
 * Get App conf
 */
type Config struct {
	Env      string `envconfig:"EVENT_ENV"`
	Port     int    `envconfig:"EVENT_PORT"`
	Password string `envconfig:"EVENT_PASSWORD,optional"`
	Stripe   struct {
		Public_Key  string `envconfig:"EVENT_STRIPE_PUBLIC,optional"`
		Private_Key string `envconfig:"EVENT_STRIPE_PRIVATE,optional"`
	}
	Email struct {
		Email    string `envconfig:"EVENT_EMAIL,optional"`
		User     string `envconfig:"EVENT_EMAIL_USER,optional"`
		Password string `envconfig:"EVENT_EMAIL_PASSWORD,optional"`
	}
	Sheet struct {
		Id string `envconfig:"EVENT_SHEET_ID,optional"`
	}
}

func Conf() *Config {
	if conf == nil {
		if err := envconfig.Init(&conf); err != nil {
			log.Fatal("err=%s\n", err)
		}
	}
	return conf
}
