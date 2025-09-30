package utils

import (
	"github.com/mdi-client-portal/client-portal-be/database/models"
)

func ToInvoiceClientResponse(invoices []models.Invoice) []models.InvoiceResponse {
	res := make([]models.InvoiceResponse, len(invoices))
	for i, c := range invoices {
		res[i] = models.InvoiceResponse{
			InvoiceID:     c.InvoiceID,
			InvoiceNumber: c.InvoiceNumber,
			DueDate:       c.DueDate,
			Total:         c.Total,
			PaymentStatus: c.PaymentStatus,
		}
	}
	return res
}
