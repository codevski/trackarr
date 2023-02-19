package main

import (
	"log"

	"github.com/codevski/trackarr/database"
	_ "github.com/codevski/trackarr/docs"
	"github.com/codevski/trackarr/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title           Trackarr API
// @version         1.0
// @description     This is the Trackarr API documentation.
// @termsOfService  https://trackarr.games/terms

// @contact.name   API Support
// @contact.url    https://discord.comn/trackarr

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Listen(":8000")
}
