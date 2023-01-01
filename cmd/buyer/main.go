package main

import (
	HealthCheckRouter "buyer/internal/healthcheck/router"
	mysql "buyer/pkg/mysql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appRouterUnversioned := app.Group("")

	db := mysql.ConnectDatabase()

	HealthCheckRouter.RegisterRoutes(appRouterUnversioned, db)

	app.Listen(":3000")
}
