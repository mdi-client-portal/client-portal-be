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
	PaymentDate     time.Time `json:"payment_date"`
	AmountPaid      float64   `json:"amount_paid"`
	VoidedAt		*time.Time `json:"voided_at"`
	ProofOfTransfer string    `json:"proof_of_transfer"`
}
