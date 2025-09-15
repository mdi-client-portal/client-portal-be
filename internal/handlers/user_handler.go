package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"github.com/mdi-client-portal/client-portal-be/internal/validators"
) 

// type NewInvoiceHandler()

func Login(c *fiber.Ctx) error {
	var req validators.UserLoginValidator

	if err := c.BodyParser(&req); err != nil{
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := validators.Validate.Struct(req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	result, err :=  

	return utils.Success(c, fiber.StatusOK, "Login success", req)
}