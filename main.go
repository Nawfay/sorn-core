package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"sorn/core/db"
)

func main() {
	app := fiber.New()

	
	db.DBInit()
	defer db.DBClose()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	SearchRoutes(app)


	app.Listen("localhost:9000")
}
