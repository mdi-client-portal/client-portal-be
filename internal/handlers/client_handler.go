package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"github.com/mdi-client-portal/client-portal-be/internal/validators"
)

type ClientHandler struct {
	service services.ClientService
}

func NewClientHandler(service services.ClientService) *ClientHandler {
	return &ClientHandler{service}
}

func (h *ClientHandler) LoginHandler(c *fiber.Ctx) error {
	var req validators.ClientLoginValidator

	if err := c.BodyParser(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := validators.Validate.Struct(req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	client, err := h.service.LoginService(req.Email, req.Password)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Login gagal", err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Login success", client)
}
