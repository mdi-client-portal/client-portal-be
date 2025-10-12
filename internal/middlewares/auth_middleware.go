package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/mdi-client-portal/client-portal-be/internal/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	fmt.Println("Authorization Header:", authHeader)

	// Periksa format header
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

	id, ok := claims["userId"].(string)
	if !ok {
		return utils.Error(c, fiber.StatusUnauthorized, "User ID not found in token", nil)
	}

	fmt.Println("User ID from token:", id)

	c.Locals("userId", id)

	return c.Next()
}
