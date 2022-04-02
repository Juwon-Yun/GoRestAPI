package routes

import (
	"github.com/Juwon-Yun/goRest/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Resister)
}
