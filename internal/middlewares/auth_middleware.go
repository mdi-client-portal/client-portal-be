package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	config.Log.Info("Auth middleware invoked")
	authHeader := c.Get("Authorization")

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		config.Log.Error("Invalid Authorization header format", zap.String("header", authHeader))
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid Authorization header format", nil)
	}
	tokenString := parts[1]

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		config.Log.Error("Failed to parse token", zap.String("token", tokenString), zap.String("error", err.Error()))	
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		config.Log.Error("Invalid token claims", zap.String("token", tokenString))
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
	}

	id, ok := claims["userId"].(string)
	if !ok {	
		return utils.Error(c, fiber.StatusUnauthorized, "User ID not found in token", nil)
	}

	config.Log.Info("User ID from token:", zap.String("user_id", id))

	c.Locals("userId", id)

	return c.Next()
}
