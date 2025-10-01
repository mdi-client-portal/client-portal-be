package utils

import (
	"github.com/mdi-client-portal/client-portal-be/database/models"
)

func ToAllPaymentByClientResponse(payments []models.Payment) []models.PaymentResponse {
	res := make([]models.PaymentResponse, len(payments))
	for i, c := range payments {
		res[i] = models.PaymentResponse{
			PaymentDate:     c.PaymentDate,
			AmountPaid:      c.AmountPaid,
			ProofOfTransfer: c.ProofOfTransfer,
		}
	}
	return res
}
