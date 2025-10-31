package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ClientService interface {
	LoginService(email string, password string) (models.ClientLoginResponse, error)
}

type clientService struct {
	repo repositories.ClientRepository
}

func NewClientService(repo repositories.ClientRepository) ClientService {
	return &clientService{repo}
}

func (cs *clientService) LoginService(email string, password string) (models.ClientLoginResponse, error) {
	config.Log.Info("Client login attempt", zap.String("email", email))

	client, err := cs.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.Log.Warn("Client login failed: email not found", zap.String("email", email))
			return models.ClientLoginResponse{}, errors.New("email not found")
		}
		return models.ClientLoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.ClientPassword), []byte(password)); err != nil {
		config.Log.Warn("Client login failed: wrong password", zap.String("email", email))
		return models.ClientLoginResponse{}, errors.New("wrong password")
	}

	config.Log.Info("Client login success", zap.String("client_id", client.ClientID))
	return utils.ToClientLoginResponse(client), nil
}
