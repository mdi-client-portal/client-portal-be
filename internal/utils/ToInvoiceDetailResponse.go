package utils

import (
	"github.com/mdi-client-portal/client-portal-be/database/models"
)

func ToInvoiceDetailResponse(c *models.Invoice) models.InvoiceDetailResponse {
	return models.InvoiceDetailResponse{
		InvoiceID:        c.InvoiceID,
		InvoiceNumber:    c.InvoiceNumber,
		IssueDate:        c.IssueDate,
		DueDate:          c.DueDate,
		TaxRate:          c.TaxRate,
		TaxAmount:        c.TaxAmount,
		SubTotal:         c.SubTotal,
		Total:            c.Total,
		TaxInvoiceNumber: c.TaxInvoiceNumber,
		AmountPaid:       c.AmountPaid,
		PaymentStatus:    c.PaymentStatus,
		ClientID:         c.ClientID,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
	}
}
