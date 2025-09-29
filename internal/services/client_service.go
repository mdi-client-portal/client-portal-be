package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
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
	client, err := cs.repo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ClientLoginResponse{}, errors.New("email tidak ditemukan")
		}
		return models.ClientLoginResponse{}, err
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(client.ClientPassword), []byte(password)); err != nil {
		return models.ClientLoginResponse{}, errors.New("password salah")
	}

	return utils.ToClientLoginResponse(client), nil
}
