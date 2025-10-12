package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	// "github.com/mdi-client-portal/client-portal-be/internal/validators"
)

type PaymentHandler struct {
	service services.PaymentService
}

func NewPaymentHandler(service services.PaymentService) *PaymentHandler {
	return &PaymentHandler{service}
}

func (h *PaymentHandler) GetAllPaymentByClientIdHandler(c *fiber.Ctx) error {
	// var req validators.PaymentClientValidator

	fmt.Println("Masuk ke get all payments handler")

	userId := c.Locals("userId").(string)

	payments, err := h.service.GetAllPaymentByClientIdService(userId)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Get Payment gagal", err.Error())
	}

	fmt.Println("Payments:", payments)
	return utils.Success(c, fiber.StatusOK, "Get Payment success", utils.ToAllPaymentByClientResponse(payments))
}
