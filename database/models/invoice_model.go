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
	InvoiceID     string `json:"invoice_id"`
	IssueDate     time.Time `json:"issue_date"`
	DueDate       time.Time `json:"due_date"`
	Total         float64 `json:"total"`
	AmountPaid    float64 `json:"amount_paid"`
	PaymentStatus string `json:"payment_status"`
}

type InvoiceDetailResponse struct {
	InvoiceID        string `json:"invoice_id"`
	InvoiceNumber    string `json:"invoice_number"`
	IssueDate        time.Time `json:"issue_date"`
	DueDate          time.Time `json:"due_date"`
	TaxRate          float64 `json:"tax_rate"`
	TaxAmount        float64 `json:"tax_amount"`
	SubTotal         float64 `json:"sub_total"`
	Total            float64 `json:"total"`
	TaxInvoiceNumber string  `json:"tax_invoice_number"`
	AmountPaid       float64 `json:"amount_paid"`
	PaymentStatus    string  `json:"payment_status"`
	ClientID         string  `json:"client_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
