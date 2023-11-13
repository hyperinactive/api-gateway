package main

import (
	"notgithub.com/hyperinactive/api-gateway/config"
	"notgithub.com/hyperinactive/api-gateway/db"
	"notgithub.com/hyperinactive/api-gateway/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := db.Connect(); err != nil {
		panic(err)
	}

	app := fiber.New()

	router.InitMiddleware(app)
	router.InitRoutes(app)

	hostUrl := config.Config.Server.Host + ":" + config.Config.Server.Port
	app.Listen(hostUrl)
}
