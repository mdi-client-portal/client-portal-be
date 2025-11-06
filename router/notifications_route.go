package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/handlers"
	"github.com/mdi-client-portal/client-portal-be/internal/middlewares"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"github.com/mdi-client-portal/client-portal-be/internal/services"
)

func NotificationRoutes(api fiber.Router) {
	notifications := api.Group("/notifications", middlewares.AuthMiddleware)

	notificationRepo := repositories.NewNotificationRepository(config.DB)
	notificationService := services.NewNotificationService(notificationRepo)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	notifications.Get("/", notificationHandler.GetAllNotificationsHandler)
	notifications.Put("/mark-as-read", notificationHandler.MarkAsReadHandler)
}
