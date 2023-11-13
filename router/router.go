package router

import (
	"notgithub.com/hyperinactive/api-gateway/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitRoutes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Mars more, sikter")
	})

	auth := app.Group("/auth")
	auth.Post("/sign-in", handler.SignIn)
	auth.Post("/sign-up", handler.SignUp)

	glupan := app.Group("/glupan")
	glupan.All("/*", func(c *fiber.Ctx) error {
		// TODO: service communication
		return c.SendStatus(fiber.StatusOK)
	})
	app.Get("/metrics", monitor.New())
}

func InitMiddleware(app *fiber.App) {
	// cors
	app.Use(cors.New())

	// log
	app.Use(logger.New())

	// recover on panic
	app.Use(recover.New())

	// caching
	app.Use(cache.New())
	app.Use(etag.New())
}
