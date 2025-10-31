package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/internal/handlers"

	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/middlewares"
	"github.com/mdi-client-portal/client-portal-be/internal/repositories"

	"github.com/mdi-client-portal/client-portal-be/internal/services"
)

func InvoiceRoutes(api fiber.Router) {
	invoice := api.Group("/invoices")

	invoiceRepo := repositories.NewInvoiceRepository(config.DB)
	invoiceService := services.NewInvoiceService(invoiceRepo)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)

	invoice.Get("/get", middlewares.AuthMiddleware, invoiceHandler.GetAllInvoiceByClientIdHandler)
	invoice.Post("/get/detail", middlewares.AuthMiddleware, invoiceHandler.GetInvoiceByIdHandler)
}
