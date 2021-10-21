package main

import (
	"github.com/bangnh1/golang-training/10/repo"
	"github.com/bangnh1/golang-training/10/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	db := repo.ConnectDatabase()
	defer db.Close()

	// Init Fiber App
	app := fiber.New()

	// Register router
	router.InitRouter(app)

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Fiber App listen port
	app.Listen(":3002")
}
