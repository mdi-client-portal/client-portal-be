package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"gorm.io/gorm"
)

type InvoiceService interface {
	GetAllInvoiceByClientIdService(clientId string) ([]models.Invoice, error)
	GetInvoiceByIdService(client string, invoiceId string) (models.InvoiceDetailResponse, error)
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

func (i *invoiceService) GetInvoiceByIdService(clientId string, invoiceId string) (models.InvoiceDetailResponse, error) {
	invoice, err := i.repo.GetInvoiceById(invoiceId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.InvoiceDetailResponse{}, errors.New("invoice tidak ditemukan")
		}
		return models.InvoiceDetailResponse{}, err
	}
	return utils.ToInvoiceDetailResponse(invoice), nil
}
