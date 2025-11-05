package utils

import (
	"time"

	"github.com/mdi-client-portal/client-portal-be/database/models"
)

func ToInvoiceClientResponse(invoices []models.Invoice) []models.InvoiceResponse {
	res := make([]models.InvoiceResponse, len(invoices))
	for i, c := range invoices {
		var voidedAt *time.Time
		if !c.VoidedAt.IsZero() {
			voidedAt = &c.VoidedAt
		}

		res[i] = models.InvoiceResponse{
			InvoiceID:     c.InvoiceID,
			InvoiceNumber: c.InvoiceNumber,
			IssueDate:     c.IssueDate,
			DueDate:       c.DueDate,
			Total:         c.Total,
			PaymentStatus: c.PaymentStatus,
			AmountPaid:    c.AmountPaid,
			VoidedAt:      voidedAt,
		}
	}
	return res
}
