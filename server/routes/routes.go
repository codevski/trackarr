package routes

import (
	"github.com/codevski/trackarr/controllers"
	_ "github.com/codevski/trackarr/docs"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Status)
	app.Get("/status", controllers.Status)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.GetUser)
}
