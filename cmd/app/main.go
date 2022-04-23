package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"quizbe/configs"
	"quizbe/configs/db"
	sentryThichlabs "quizbe/configs/sentry"
	"quizbe/routes"
)

func main() {
	// Custom config
	app := fiber.New(fiber.Config{
		ServerHeader: configs.Configs.AppName,
		AppName:      configs.Configs.AppName,
		BodyLimit:    100 * 1024,
	})
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Load configs
	configs.LoadConfig()
	sentryThichlabs.InitInstanceSentry()
	db.ConnectPostgres()
	//redis.CreateRedisClient()

	// Implement Compress
	app.Use(compress.New())
	// Implement CORS
	app.Use(cors.New())
	// Implement ETAG
	app.Use(etag.New())
	app.Use(favicon.New())

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			switch e.(type) {
			case error:
				sentry.CaptureException(e.(error))
			case string:
				sentry.CaptureMessage(e.(string))
			}
		},
	}))

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")
		c.Set("Content-Type", "application/json")

		return c.Next()
	})

	routes.InitRoutes(app)

	log.Fatal(app.Listen(":" + configs.Configs.Port))
}
