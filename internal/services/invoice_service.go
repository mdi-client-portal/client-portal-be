package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"gorm.io/gorm"
)

type InvoiceService interface {
	GetAllInvoiceByClientIdService(clientId string) ([]models.Invoice, error)
	GetInvoiceByIdService(invoiceId string) (*models.InvoiceWithDetailResponse, error)
}

type invoiceService struct {
	repo repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) InvoiceService {
	return &invoiceService{repo}
}

func (i *invoiceService) GetAllInvoiceByClientIdService(clientId string) ([]models.Invoice, error) {
	invoices, err := i.repo.GetAllInvoiceByClientId(clientId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invoice tidak ditemukan")
		}
		return nil, err
	}

	return invoices, nil
}

func (i *invoiceService) GetInvoiceByIdService(invoiceId string) (*models.InvoiceWithDetailResponse, error) {
	return i.repo.GetInvoiceById(invoiceId)
}