package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"github.com/mdi-client-portal/client-portal-be/internal/validators"
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service}
}

func (h *PaymentHandler) GetAllPaymentByClientIdHandler(c *fiber.Ctx) error {
	var req validators.PaymentClientValidator

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := validators.Validate.Struct(req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	payments, err := h.service.GetAllPaymentByClientIdService(req.ClientId)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Get Payment gagal", err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Get Payment success", utils.ToAllPaymentByClientResponse(payments))
}
