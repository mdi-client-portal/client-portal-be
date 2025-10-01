package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
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
	payments, err := p.repo.GetAllPaymentByClientId(clientId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment tidak ditemukan")
		}
		return nil, err
	}

	return payments, nil
}
