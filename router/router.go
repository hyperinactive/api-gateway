package router

import (
	"microservice/handler"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Mars more, sikter")
	})

	auth := app.Group("/auth")
	auth.Post("/sign-in", handler.SignIn)
	auth.Post("/sign-up", handler.SignUp)
}
