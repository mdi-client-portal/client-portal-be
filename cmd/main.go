package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/router"
)

func main() {
    cfg := config.Load()

	// config.ConnectDB(cfg)

    // seeders.UserSeeder(config.DB)
	// seeders.ClientSeeder(config.DB)
	// seeders.InvoiceSeeder(config.DB)

    app := fiber.New()
	router.SetupUserRoutes(app)
    app.Listen("0.0.0.0:" + cfg.Port)
}