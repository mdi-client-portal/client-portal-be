package models

import "time"

type Invoice struct {
	InvoiceID        string
	InvoiceNumber    string
	IssueDate        time.Time
	DueDate          time.Time
	TaxRate          float64
	TaxAmount        float64
	SubTotal         float64
	Total            float64
	TaxInvoiceNumber string
	AmountPaid       float64
	PaymentStatus    string
	ClientID         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	VoidedAt         time.Time
}

type InvoiceResponse struct {
	InvoiceID     string
	InvoiceNumber string
	DueDate       time.Time
	Total         float64
	PaymentStatus string
}

type InvoiceDetailResponse struct {
	InvoiceID        string
	InvoiceNumber    string
	IssueDate        time.Time
	DueDate          time.Time
	TaxRate          float64
	TaxAmount        float64
	SubTotal         float64
	Total            float64
	TaxInvoiceNumber string
	AmountPaid       float64
	PaymentStatus    string
	ClientID         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
