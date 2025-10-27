package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/internal/jobs"

	"github.com/robfig/cron/v3"
	// "github.com/mdi-client-portal/client-portal-be/database/seeders"
	"github.com/mdi-client-portal/client-portal-be/router"
)

func main() {
	cfg := config.Load()

	config.ConnectDB(cfg)

	// seeders.ClientSeeder(config.DB)
	// seeders.InvoiceSeeder(config.DB)

	cronJob := cron.New()
	cronJob.AddFunc("* * * * *", func() {jobs.EmailCron(config.DB)})
	cronJob.Start()

	app := fiber.New()
	app.Use(cors.New())
	router.SetupUserRoutes(app)
	app.Listen("0.0.0.0:" + cfg.Port)
}
