package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/mdi-client-portal/client-portal-be/config"
	"github.com/mdi-client-portal/client-portal-be/database/seeders"
	"github.com/mdi-client-portal/client-portal-be/internal/jobs"
	"github.com/mdi-client-portal/client-portal-be/internal/utils"
	"github.com/mdi-client-portal/client-portal-be/router"
	"github.com/robfig/cron/v3"
)

func main() {
	config.LoggerInit()
	config.EnvInit()

	config.Log.Info("Starting application...")

	defer config.Log.Sync()

	config.ConnectDB()

	// Run seeders
	// seeders.ClientSeeder(config.DB)
	// seeders.InvoiceSeeder(config.DB)
	seeders.PaymentSeeder(config.DB)

	config.Log.Info("Setting up cron jobs...")
	cronJob := cron.New()
	cronJob.AddFunc("0 7 * * *", func() { jobs.EmailCron(config.DB) })
	cronJob.Start()
	config.Log.Info("Cron jobs set up successfully")

	config.Log.Info("Setting up routes and starting server...")
	app := fiber.New()
	app.Use(cors.New())

	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return utils.Error(c, fiber.StatusTooManyRequests, "Too Many Request", "Please try again later")
		},
	}))

	router.SetupRoutes(app)

	config.Log.Info("Server is running on port " + config.Env.Port)
	app.Listen("0.0.0.0:" + config.Env.Port)
}
