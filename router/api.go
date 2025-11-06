package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router *fiber.App) {
	api := router.Group("/api")

	ClientRoutes(api)
	InvoiceRoutes(api)
	PaymentRoutes(api)
	NotificationRoutes(api)
}
