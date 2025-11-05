package repositories

import (
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetAllPaymentByClientId(clientId string) ([]models.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (p *paymentRepository) GetAllPaymentByClientId(clientId string) ([]models.Payment, error) {
	var payments []models.Payment
	query := `
		SELECT 
			p.payment_date,
			p.amount_paid,
			p.voided_at,
			p.proof_of_transfer,
			i.invoice_number
		FROM payments p
		JOIN invoices i 
			ON p.invoice_id = i.invoice_id
		JOIN clients c 
			ON i.client_id = c.client_id
		WHERE c.client_id = ?
		ORDER BY p.payment_date DESC;

	`
	if err := p.db.Raw(query, clientId).Scan(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
