package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PaymentService interface {
	GetAllPaymentByClientIdService(clientId string) ([]models.Payment, error)
}

type paymentService struct {
	repo repositories.PaymentRepository
}

func NewPaymentService(repo repositories.PaymentRepository) PaymentService {
	return &paymentService{repo}
}

func (p *paymentService) GetAllPaymentByClientIdService(clientId string) ([]models.Payment, error) {
	config.Log.Info("Get all payments attempt", zap.String("client_id", clientId))
	
	payments, err := p.repo.GetAllPaymentByClientId(clientId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.Log.Warn("Get all payments failed: no payments found", zap.String("client_id", clientId))
			return nil, errors.New("payment tidak ditemukan")
		}

		config.Log.Error("Get all payments failed: ", zap.String("error", err.Error()))
		return nil, err
	}

	config.Log.Info("Get all payments success", zap.String("client_id", clientId))
	return payments, nil
}
