package routes

import (
	"github.com/gofiber/fiber/v2"
	"quizbe/controllers"
)

func AppRoutes(app fiber.Router) {
	app.Post("/app", controllers.CreateApp)
	app.Get("/apps", controllers.GetApps)
	app.Get("/app", controllers.GetApp)
	app.Patch("/app", controllers.UpdateApp)
	app.Delete("/app", controllers.DeleteApp)
}
