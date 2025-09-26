package handlers

import (
	"fmt"

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
	fmt.Println(req.Email)
	fmt.Println(req.Password )

	if(req.Email != "user@example.com" || req.Password != "password"){
		return utils.Error(c, fiber.StatusNotFound, "Login failed", nil)
	}

	return utils.Success(c, fiber.StatusOK, "Login success", fiber.Map{
		"id":    "1",
		"email": req.Email,
		"name":  "User",
	})
}