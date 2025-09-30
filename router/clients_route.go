package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/handlers"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
)

func ClientRoutes(api fiber.Router) {
	client := api.Group("/clients")

	clientRepo := repositories.NewClientRepository(config.DB)
	clientService := services.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(clientService)

	client.Post("/login", clientHandler.LoginHandler)
}
