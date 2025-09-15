package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/handlers"
)

func UserRoutes(api fiber.Router) {
	user := api.Group("/user")

	user.Post("/login", handlers.Login)
	// user.Post("/register", controllers.UserRegister)
}