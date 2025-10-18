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

type InvoiceDetail struct {
	InvoiceDetailID  string
	InvoiceID        string
	Amount           float64
	CreatedAt        time.Time
	PricePerDelivery float64
	TransactionNote  string
	UpdatedAt        time.Time
	DeliveryCount    int
	DeletedAt        time.Time
}


type InvoiceResponse struct {
	InvoiceID     string `json:"invoice_id"`
	IssueDate     time.Time `json:"issue_date"`
	DueDate       time.Time `json:"due_date"`
	Total         float64 `json:"total"`
	AmountPaid    float64 `json:"amount_paid"`
	PaymentStatus string `json:"payment_status"`
}

type InvoiceExtendedResponse struct {
	InvoiceID        string `json:"invoice_id"`
	InvoiceNumber    string `json:"invoice_number"`
	IssueDate        time.Time `json:"issue_date"`
	DueDate          time.Time `json:"due_date"`
	TaxRate          float64   `json:"tax_rate"`
	TaxAmount        float64   `json:"tax_amount"`
	SubTotal         float64   `json:"sub_total"`
	Total            float64   `json:"total"`
	TaxInvoiceNumber string    `json:"tax_invoice_number"`
	AmountPaid       float64   `json:"amount_paid"`
	PaymentStatus    string    `json:"payment_status"`
	VoidedAt         time.Time `json:"voided_at"`
}

type InvoiceDetailResponse struct {
	InvoiceDetailID   string    `json:"invoice_detail_id"`
	InvoiceID         string    `json:"invoice_id"`
	Amount            float64   `json:"amount"`
	CreatedAt         time.Time `json:"created_at"`
	PricePerDelivery  float64   `json:"price_per_delivery"`
	TransactionNote   string    `json:"transaction_note"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeliveryCount     int       `json:"delivery_count"`
	DeletedAt         time.Time `json:"deleted_at"`
}

type InvoiceWithDetailResponse struct {
	Invoice        InvoiceExtendedResponse  `json:"invoice"`
	InvoiceDetails []InvoiceDetailResponse `json:"invoice_details"`
}