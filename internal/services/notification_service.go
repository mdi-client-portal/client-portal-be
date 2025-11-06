package services

import (
	"errors"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/database/models"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NotificationService interface {
	GetAllNotificationByClientIdService(clientId string) ([]models.Notification, error)
	MarkAsReadService(clientId string) error
	CreateNotificationService(clientId, message string) error
}

type notificationService struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationService{repo}
}

func (n *notificationService) GetAllNotificationByClientIdService(clientId string) ([]models.Notification, error) {
	config.Log.Info("Get all notifications attempt", zap.String("client_id", clientId))
	
	notifications, err := n.repo.GetAllNotificationByClientId(clientId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.Log.Warn("Get all notifications failed: no notifications found", zap.String("client_id", clientId))
			return nil, errors.New("notification tidak ditemukan")
		}
		config.Log.Error("Get all notifications failed: ", zap.String("error", err.Error()))
		return nil, err
	}

	config.Log.Info("Get all notifications success", zap.String("client_id", clientId))
	return notifications, nil
}

func (n *notificationService) MarkAsReadService(clientId string) error {
	config.Log.Info("Mark as read attempt", zap.String("client_id", clientId))
	
	if err := n.repo.MarkAsRead(clientId); err != nil {
		config.Log.Error("Mark as read failed: ", zap.String("error", err.Error()))
		return err
	}

	config.Log.Info("Mark as read success", zap.String("client_id", clientId))
	return nil
}

func (n *notificationService) CreateNotificationService(clientId, message string) error {
	config.Log.Info("Create notification attempt", zap.String("client_id", clientId), zap.String("message", message))
	
	if err := n.repo.CreateNotification(clientId, message); err != nil {
		config.Log.Error("Create notification failed: ", zap.String("error", err.Error()))
		return err
	}

	config.Log.Info("Create notification success", zap.String("client_id", clientId))
	return nil
}
