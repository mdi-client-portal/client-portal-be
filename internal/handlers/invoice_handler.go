package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	// "github.com/mdi-client-portal/client-portal-be/internal/validators"
)

type InvoiceHandler struct {
	service services.InvoiceService
}

func NewInvoiceHandler(service services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service}
}

func (h *InvoiceHandler) GetAllInvoiceByClientIdHandler(c *fiber.Ctx) error {
	fmt.Println("Masuk ke get all invoice handler") 

	authHeader := c.Get("Authorization")
	fmt.Println("Authorization Header:", authHeader)
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid Authorization header format", nil)
	}
	tokenString := parts[1]

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
	}

	fmt.Println("Decoded claims:", claims)

	id, ok := claims["userId"].(string)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "ID not found in token", nil)
	}

	fmt.Println("User ID:", id)

	invoices, err := h.service.GetAllInvoiceByClientIdService(id)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Get Invoice gagal", err.Error())
	}

	return utils.Success(c, fiber.StatusOK, "Get Invoice success", utils.ToInvoiceClientResponse(invoices))
}

func (h *InvoiceHandler) GetInvoiceByIdHandler(c *fiber.Ctx) error {
	type RequestBody struct {
		InvoiceID string `json:"invoice_id"`
	}

	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if body.InvoiceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invoice_id is required",
		})
	}

	response, err := h.service.GetInvoiceByIdService(body.InvoiceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}