package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppName string `mapstructure:"APP_NAME"`
	AppUrl  string `mapstructure:"APP_URL"`
	Port    string `mapstructure:"PORT"`

	DBUrl string `mapstructure:"DATABASE_URL"`

	EmailFrom    string `mapstructure:"EMAIL_FROM"`
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	SmtpUsername string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`

	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func LoadEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &env
}
