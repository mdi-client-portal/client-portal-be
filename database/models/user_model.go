package models

import "time"

type User struct {
	ClientID      string	`gorm:"primaryKey"`
	Currency      string	`gorm:"size:100;not null"`
	Email         string	`gorm:"size:150;uniqueIndex;not null"`
	Password      string	`gorm:"size:255;not null"`
	Country       string	`gorm:"size:100"`
	ClientAddress string	`gorm:"size:255"`
	PostalCode    string	`gorm:"size:20"`
	ClientPhone   string	`gorm:"size:20"`
	DeletedAt     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}