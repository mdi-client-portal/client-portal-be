package handlers

import (
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	
	
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service}
}

func (h *PaymentHandler) GetAllPaymentByClientIdHandler(c *fiber.Ctx) error {
	config.Log.Info("Get all payments attempt")

	userId := c.Locals("userId").(string)

	payments, err := h.service.GetAllPaymentByClientIdService(userId)
	if err != nil {
		config.Log.Error("Failed to get all payments: ", zap.String("error", err.Error()))
		return utils.Error(c, fiber.StatusUnauthorized, "Get Payment gagal", err.Error())
	}

	config.Log.Info("Get all payments success", zap.String("client_id", userId))
	return utils.Success(c, fiber.StatusOK, "Get Payment success", utils.ToAllPaymentByClientResponse(payments))
}
