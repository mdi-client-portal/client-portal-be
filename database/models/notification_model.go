package models

import "time"

type Notification struct {
	NotificationID int
	ClientID       string
	Message        string
	Read           bool
	CreatedAt      time.Time
}

type NotificationResponse struct {
	NotificationID int       `json:"notification_id"`
	ClientID       string    `json:"client_id"`
	Message        string    `json:"message"`
	Read           bool      `json:"read"`
	CreatedAt      time.Time `json:"created_at"`
}