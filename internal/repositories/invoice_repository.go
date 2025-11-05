package repositories

import (
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	GetAllInvoiceByClientId(clientId string) ([]models.Invoice, error)
	GetInvoiceById(invoiceId string) (*models.InvoiceWithDetailResponse, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db}
}

func (i *invoiceRepository) GetAllInvoiceByClientId(clientId string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	query := `
		SELECT *
		FROM invoices WHERE client_id = ?
	`
	if err := i.db.Raw(query, clientId).Scan(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *invoiceRepository) GetInvoiceById(invoiceId string) (*models.InvoiceWithDetailResponse, error) {
	var invoice models.Invoice
	var invoiceDetails []models.InvoiceDetail

	if err := r.db.Raw(`SELECT * FROM invoices WHERE invoice_id = ? LIMIT 1`, invoiceId).Scan(&invoice).Error; err != nil {
		return nil, err
	}

	// Check if invoice was found
	if invoice.InvoiceID == "" {
		return nil, gorm.ErrRecordNotFound
	}

	if err := r.db.Raw(`SELECT * FROM invoice_details WHERE invoice_id = ?`, invoiceId).Scan(&invoiceDetails).Error; err != nil {
		return nil, err
	}

	response := &models.InvoiceWithDetailResponse{
		Invoice: models.InvoiceExtendedResponse{
			InvoiceID:        invoice.InvoiceID,
			InvoiceNumber:    invoice.InvoiceNumber,
			IssueDate:        invoice.IssueDate,
			DueDate:          invoice.DueDate,
			TaxRate:          invoice.TaxRate,
			TaxAmount:        invoice.TaxAmount,
			SubTotal:         invoice.SubTotal,
			Total:            invoice.Total,
			TaxInvoiceNumber: invoice.TaxInvoiceNumber,
			AmountPaid:       invoice.AmountPaid,
			PaymentStatus:    invoice.PaymentStatus,
			VoidedAt:         invoice.VoidedAt,
		},
	}

	for _, d := range invoiceDetails {
		response.InvoiceDetails = append(response.InvoiceDetails, models.InvoiceDetailResponse{
			InvoiceDetailID:  d.InvoiceDetailID,
			InvoiceID:        d.InvoiceID,
			Amount:           d.Amount,
			CreatedAt:        d.CreatedAt,
			PricePerDelivery: d.PricePerDelivery,
			TransactionNote:  d.TransactionNote,
			UpdatedAt:        d.UpdatedAt,
			DeliveryCount:    d.DeliveryCount,
			DeletedAt:        d.DeletedAt,
		})
	}

	return response, nil
}

func GetUnpaidAndPartialInvoices(db *gorm.DB) ([]models.Invoice, error) {
	var invoices []models.Invoice
	if err := db.Where("payment_status IN (?, ?)", "unpaid", "partial").Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

//