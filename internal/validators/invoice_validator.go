package validators

type InvoiceClientValidator struct {
	ClientId string `json:"client_id" validate:"required"`
}

type InvoiceDetailValidator struct {
	ClientId  string `json:"client_id" validate:"required"`
	InvoiceId string `json:"invoice_id" validate:"required"`
}
