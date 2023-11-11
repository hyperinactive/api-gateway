package main

import (
	"microservice/config"
	"microservice/db"
	"microservice/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := db.Connect(); err != nil {
		panic(err)
	}

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

	app.Get("/metrics", monitor.New())

	router.Init(app)

	hostUrl := config.Config.Server.Host + ":" + config.Config.Server.Port
	app.Listen(hostUrl)
}
