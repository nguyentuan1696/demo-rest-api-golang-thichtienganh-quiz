package routes

import (
	"github.com/gofiber/fiber/v2"
	//"quizbe/middlewares"
)

func InitRoutes(app *fiber.App) {
	QuizRoutes(app.Group("/v1"))
	AppRoutes(app.Group("/v1"))
}
