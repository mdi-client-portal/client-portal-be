package handlers

import (
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
)

type NotificationHandler struct {
	service services.NotificationService
}

func NewNotificationHandler(service services.NotificationService) *NotificationHandler {
	return &NotificationHandler{service}
}

func (h *NotificationHandler) GetAllNotificationsHandler(c *fiber.Ctx) error {
	config.Log.Info("Get all notifications attempt")

	clientId := c.Locals("userId").(string)

	notifications, err := h.service.GetAllNotificationByClientIdService(clientId)
	if err != nil {
		config.Log.Error("Failed to get all notifications: ", zap.String("error", err.Error()))
		return utils.Success(c, fiber.StatusOK, "Get Notifications success", []interface{}{})
	}

	config.Log.Info("Get all notifications success", zap.String("client_id", clientId))
	return utils.Success(c, fiber.StatusOK, "Get Notifications success", notifications)
}

func (h *NotificationHandler) MarkAsReadHandler(c *fiber.Ctx) error {
	config.Log.Info("Mark as read attempt")

	clientId := c.Locals("userId").(string)

	if err := h.service.MarkAsReadService(clientId); err != nil {
		config.Log.Error("Failed to mark as read: ", zap.String("error", err.Error()))
		return utils.Error(c, fiber.StatusInternalServerError, "Mark as read gagal", err.Error())
	}

	config.Log.Info("Mark as read success", zap.String("client_id", clientId))
	return utils.Success(c, fiber.StatusOK, "Mark as read success", nil)
}
