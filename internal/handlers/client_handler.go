package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/mdi-client-portal/client-portal-be/config"
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
	config.Log.Info("Client login attempt")

	var req validators.ClientLoginValidator

	if err := c.BodyParser(&req); err != nil {
		config.Log.Error("Failed to parse request body: ", zap.String("error", err.Error()))
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := validators.Validate.Struct(req); err != nil {
		config.Log.Error("Validation failed: ", zap.String("error", err.Error()))
		return utils.Error(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	client, err := h.service.LoginService(req.Email, req.Password)
	if err != nil {
		config.Log.Error("Login failed: ", zap.String("error", err.Error()))
		return utils.Error(c, fiber.StatusUnauthorized, "Login failed", err.Error())
	}

	config.Log.Info("Client logged in successfully", zap.String("client_email", req.Email))
	return utils.Success(c, fiber.StatusOK, "Login success", client)
}
