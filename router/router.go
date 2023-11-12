package router

import (
	"notgithub.com/hyperinactive/api-gateway/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Init(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Mars more, sikter")
	})

	auth := app.Group("/auth")
	auth.Post("/sign-in", handler.SignIn)
	auth.Post("/sign-up", handler.SignUp)

	app.Get("/metrics", monitor.New())
}
