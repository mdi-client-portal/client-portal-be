package repositories

import (
	"github.com/mdi-client-portal/client-portal-be/database/models"
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
	var client models.Client
	query := `
		SELECT *
		FROM clients WHERE client_email = ? LIMIT 1
	`
	if err := r.db.Raw(query, email).Scan(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}
