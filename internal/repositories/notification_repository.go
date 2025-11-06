package repositories

import (
	"time"

	"github.com/mdi-client-portal/client-portal-be/database/models"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	GetAllNotificationByClientId(clientId string) ([]models.Notification, error)
	MarkAsRead(clientId string) error
	CreateNotification(clientId, message string) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (n *notificationRepository) GetAllNotificationByClientId(clientId string) ([]models.Notification, error) {
	var notifications []models.Notification
	
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	
	query := `
		SELECT 
			notification_id,
			client_id,
			message,
			read,
			created_at
		FROM notifications
		WHERE client_id = ? AND created_at >= ?
		ORDER BY created_at DESC
	`
	
	if err := n.db.Raw(query, clientId, sevenDaysAgo).Scan(&notifications).Error; err != nil {
		return nil, err
	}
	
	return notifications, nil
}

func (n *notificationRepository) MarkAsRead(clientId string) error {
	query := `
		UPDATE notifications
		SET read = true
		WHERE client_id = ? AND read = false
	`
	
	if err := n.db.Exec(query, clientId).Error; err != nil {
		return err
	}
	
	return nil
}

func (n *notificationRepository) CreateNotification(clientId, message string) error {
	query := `INSERT INTO notifications (client_id, message, read, created_at) 
		VALUES (?, ?, ?, ?)`

	if err := n.db.Exec(query, clientId, message, false, time.Now()).Error; err != nil {
		return err
	}

	return nil
}
