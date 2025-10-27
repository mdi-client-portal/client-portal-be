package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
	DBHost string
	DBName string
	DBUser string
	DBPassword string
	DBPort string
	FromEmail string
	FromEmailPassword string
	FromEmailSMTP string
	SMTPAddr string
}

func Load() *Config{
	v := viper.New()

	v.SetDefault("PORT", "3000")
	v.SetDefault("DB_HOST", "localhost")
	v.SetDefault("DB_NAME", "APP_DB")
	v.SetDefault("DB_USER", "postgres")
	v.SetDefault("DB_PASSWORD", "12345")
	v.SetDefault("DB_PORT", "5432")
	v.SetDefault("FROM_EMAIL", "example@gmail.com")
	v.SetDefault("FROM_EMAIL_PASSWORD", "xxxx xxxx xxxx xxxx")
	v.SetDefault("FROM_EMAIL_SMTP", "smtp.gmail.com")
	v.SetDefault("SMTP_ADDR", "smtp.gmail.com:587")

	v.SetConfigFile(".env")
    v.AutomaticEnv() 

	if err := v.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables only")
	}

	return &Config{
        Port: v.GetString("PORT"),
		DBHost : v.GetString("DB_HOST"),
		DBName : v.GetString("DB_NAME"),
		DBUser : v.GetString("DB_USER"),
		DBPassword : v.GetString("DB_PASSWORD"),
		DBPort : v.GetString("DB_PORT"),
		FromEmail: v.GetString("FROM_EMAIL"),
		FromEmailPassword: v.GetString("FROM_EMAIL_PASSWORD"),
		FromEmailSMTP: v.GetString("FROM_EMAIL_SMTP"),
		SMTPAddr: v.GetString("SMTP_ADDR"),
    }
    
}