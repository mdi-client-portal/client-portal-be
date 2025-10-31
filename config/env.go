package config

import (
	"github.com/spf13/viper"
)

type Environment struct {
	Port              string
	DBHost            string
	DBName            string
	DBUser            string
	DBPassword        string
	DBPort            string
	FromEmail         string
	FromEmailPassword string
	FromEmailSMTP     string
	SMTPAddr          string
	AppEnv            string
	LogLevel          string
}

var Env *Environment

func EnvInit() *Environment {
	Log.Info("Initializing environment variables...")
	v := viper.New()

	v.SetDefault("PORT", "3000")
	v.SetDefault("DB_HOST", "localhost")
	v.SetDefault("DB_NAME", "APP_DB")
	v.SetDefault("DB_USER", "postgres")
	v.SetDefault("DB_PASSWORD", "12345")
	v.SetDefault("DB_PORT", "5432")
	v.SetDefault("APP_ENV", "development")
	v.SetDefault("LOG_LEVEL", "info")

	v.SetConfigFile(".env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		Log.Warn("No .env file found, reading configuration from environment variables")
	}

	Env = &Environment{
		Port:              v.GetString("PORT"),
		DBHost:            v.GetString("DB_HOST"),
		DBName:            v.GetString("DB_NAME"),
		DBUser:            v.GetString("DB_USER"),
		DBPassword:        v.GetString("DB_PASSWORD"),
		DBPort:            v.GetString("DB_PORT"),
		FromEmail:         v.GetString("FROM_EMAIL"),
		FromEmailPassword: v.GetString("FROM_EMAIL_PASSWORD"),
		FromEmailSMTP:     v.GetString("FROM_EMAIL_SMTP"),
		SMTPAddr:          v.GetString("SMTP_ADDR"),
		AppEnv:            v.GetString("APP_ENV"),
		LogLevel:          v.GetString("LOG_LEVEL"),
	}

	Log.Info("Environment variables initialized")
	return Env
}
