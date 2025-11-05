package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"go.uber.org/zap"
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
	config.Log.Info("Get all invoices attempt", zap.String("client_id", clientId))
	
	invoices, err := i.repo.GetAllInvoiceByClientId(clientId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.Log.Warn("Get all invoices failed: no invoices found", zap.String("client_id", clientId))
			return nil, errors.New("invoice tidak ditemukan")
		}
		return nil, err
	}

	config.Log.Info("Get all invoices success", zap.String("client_id", clientId))
	return invoices, nil
}

func (i *invoiceService) GetInvoiceByIdService(invoiceId string) (*models.InvoiceWithDetailResponse, error) {
	config.Log.Info("Get invoice by ID attempt", zap.String("invoice_id", invoiceId))
	
	response, err := i.repo.GetInvoiceById(invoiceId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.Log.Warn("Get invoice by ID failed: invoice not found", zap.String("invoice_id", invoiceId))
			return nil, errors.New("invoice tidak ditemukan")
		}
		config.Log.Error("Get invoice by ID failed: ", zap.String("error", err.Error()))
		return nil, err
	}

	config.Log.Info("Get invoice by ID success", zap.String("invoice_id", invoiceId))
	return response, nil
}