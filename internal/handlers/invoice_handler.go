package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"github.com/mdi-client-portal/client-portal-be/internal/validators"
)

type InvoiceHandler struct {
	service services.InvoiceService
}

func NewInvoiceHandler(service services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service}
}

func (h *InvoiceHandler) GetAllInvoiceByClientIdHandler(c *fiber.Ctx) error {
	var req validators.InvoiceClientValidator

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := validators.Validate.Struct(req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	invoices, err := h.service.GetAllInvoiceByClientIdService(req.ClientId)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Get Invoice gagal", err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Get Invoice success", utils.ToInvoiceClientResponse(invoices))
}

func (h *InvoiceHandler) GetInvoiceByIdHandler(c *fiber.Ctx) error {
	var req validators.InvoiceDetailValidator

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := validators.Validate.Struct(req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	client, err := h.service.GetInvoiceByIdService(req.ClientId, req.InvoiceId)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Get Invoice gagal", err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Get Invoice success", client)
}
