package utils

import (
	"time"

	"github.com/mdi-client-portal/client-portal-be/database/models"
)

func ToAllPaymentByClientResponse(payments []models.Payment) []models.PaymentResponse {
	res := make([]models.PaymentResponse, len(payments))

	for i, c := range payments {
		var voidedAt *time.Time
		if !c.VoidedAt.IsZero() {
			voidedAt = &c.VoidedAt
		}

		res[i] = models.PaymentResponse{
			PaymentDate:     c.PaymentDate,
			AmountPaid:      c.AmountPaid,
			VoidedAt:        voidedAt,
			ProofOfTransfer: c.ProofOfTransfer,
			InvoiceNumber:   c.InvoiceNumber,
		}
	}

	return res
}