package repositories

import (
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ClientRepository interface {
	FindByEmail(email string) (*models.Client, error)
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db}
}

func (r *clientRepository) FindByEmail(email string) (*models.Client, error) {
	config.Log.Info("Find client by email", zap.String("email", email))

	var client models.Client
	query := `
		SELECT *
		FROM clients WHERE client_email = ? LIMIT 1
	`
	if err := r.db.Raw(query, email).Scan(&client).Error; err != nil {
		config.Log.Error("Find client by email failed", zap.String("email", email), zap.Error(err))
		return nil, err
	}
	return &client, nil
}
 
func (r *clientRepository) GetEmailByClientId(clientId string) (string, error) {
	config.Log.Info("Get email by client ID", zap.String("client_id", clientId))

	var client models.Client
	query := `
		SELECT *
		FROM clients WHERE client_id = ? LIMIT 1
	`
	if err := r.db.Raw(query, clientId).Scan(&client).Error; err != nil {
		config.Log.Error("Get email by client ID failed", zap.String("client_id", clientId), zap.Error(err))
		return "", err
	}

	config.Log.Info("Get email by client ID success", zap.String("client_id", clientId))
	return client.ClientEmail, nil
}