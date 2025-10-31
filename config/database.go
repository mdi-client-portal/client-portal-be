package config

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {	
	Log.Info("Connecting to database...")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		Env.DBHost,
		Env.DBUser,
		Env.DBPassword,
		Env.DBName,
		Env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Error("failed to connect database: ", zap.String("error", err.Error()))
	}

	DB = db
	Log.Info("Database connected")
}