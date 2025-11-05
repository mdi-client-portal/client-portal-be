package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"go.uber.org/zap"
)

type InvoiceHandler struct {
	service services.InvoiceService
}

func NewInvoiceHandler(service services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service}
}

func (h *InvoiceHandler) GetAllInvoiceByClientIdHandler(c *fiber.Ctx) error {
	config.Log.Info("Get all invoices attempt")

	userId := c.Locals("userId").(string)

	invoices, err := h.service.GetAllInvoiceByClientIdService(userId)
	if err != nil {
		config.Log.Error("Failed to get all invoices: ", zap.String("error", err.Error()))
		return utils.Error(c, fiber.StatusUnauthorized, "Get Invoice gagal", err.Error())
	}

	config.Log.Info("Get all invoices success", zap.String("client_id", userId))
	return utils.Success(c, fiber.StatusOK, "Get Invoice success", utils.ToInvoiceClientResponse(invoices))
}

func (h *InvoiceHandler) GetInvoiceByIdHandler(c *fiber.Ctx) error {
	config.Log.Info("Get invoice by ID attempt")

	type RequestBody struct {
		InvoiceID string `json:"invoice_id"`
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		config.Log.Error("Failed to parse request body: ", zap.String("error", err.Error()))
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if body.InvoiceID == "" {
		config.Log.Error("invoice_id is required")
		return utils.Error(c, fiber.StatusBadRequest, "invoice_id is required", nil)
	}

	response, err := h.service.GetInvoiceByIdService(body.InvoiceID)
	if err != nil {
		config.Log.Error("Failed to get invoice by ID: ", zap.String("error", err.Error()))
		return utils.Success(c, fiber.StatusOK, "not found", nil)
	}

	config.Log.Info("Get invoice by ID success", zap.String("invoice_id", body.InvoiceID))
	return utils.Success(c, fiber.StatusOK, "Get Invoice by ID success", response)
}