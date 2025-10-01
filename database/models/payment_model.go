package models

import "time"

type Payment struct {
	PaymentID       string
	PaymentDate     time.Time
	AmountPaid      float64
	ProofOfTransfer string
	InvoiceID       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	VoidedAt        time.Time
	InvoiceNumber   string
}

type PaymentResponse struct {
	PaymentDate     time.Time
	AmountPaid      float64
	ProofOfTransfer string
}
