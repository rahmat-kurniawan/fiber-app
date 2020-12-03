package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahmat-kurniawan/fiber-app/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", controllers.HelloWorld)
	v1.Get("/employee/", controllers.ListEmployee)
	v1.Get("/employee/:id", controllers.ShowEmployee)
}
