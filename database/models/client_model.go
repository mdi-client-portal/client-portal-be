package models

import "time"

type Client struct {
	ClientID       string
	ClientName     string
	Currency       string
	Country        string
	ClientAddress  string
	PostalCode     string
	ClientEmail    string
	ClientPassword string
	DeletedAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ClientResponse struct {
	ClientID      string    `json:"client_id"`
	ClientEmail   string    `json:"client_email"`
	ClientName    string    `json:"client_name"`
	Currency      string    `json:"currency"`
	Country       string    `json:"country"`
	ClientAddress string    `json:"client_address"`
	PostalCode    string    `json:"postal_code"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ClientLoginResponse struct {
	ClientID    string `json:"client_id"`
	ClientName  string `json:"client_name"`
	ClientEmail string `json:"client_email"`
}
