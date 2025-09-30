package repositories

import (
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	GetAllInvoiceByClientId(clientId string) ([]models.Invoice, error)
	GetInvoiceById(invoiceId string) (*models.Invoice, error)
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

func (i *invoiceRepository) GetInvoiceById(invoiceId string) (*models.Invoice, error) {
	var invoice models.Invoice
	query := `
		SELECT *
		FROM invoices WHERE invoice_id = ? LIMIT 1
	`
	if err := i.db.Raw(query, invoiceId).Scan(&invoice).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}
