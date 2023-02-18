package main

import (
	"log"

	"github.com/codevski/trackarr/database"
	"github.com/codevski/trackarr/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, // Allow cookies to be sent to the FE
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
