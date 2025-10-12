package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/handlers"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"
	"github.com/mdi-client-portal/client-portal-be/internal/services"

	"github.com/mdi-client-portal/client-portal-be/internal/middlewares"
)

func PaymentRoutes(api fiber.Router) {
	payment := api.Group("/payments")

	paymentRepo := repositories.NewPaymentRepository(config.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	payment.Get("/get", middlewares.AuthMiddleware,paymentHandler.GetAllPaymentByClientIdHandler)
}
