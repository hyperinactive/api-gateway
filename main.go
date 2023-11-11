package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// cors
	app.Use(cors.New())

	// log
	app.Use(logger.New())

	// recover on panic
	app.Use(recover.New())

	// caching
	app.Use(cache.New())
	app.Use(etag.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Mars more, sikter")
	})

	app.Get("/metrics", monitor.New())

	app.Listen("127.0.0.1:8080")
}
