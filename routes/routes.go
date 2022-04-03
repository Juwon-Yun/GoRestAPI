package routes

import (
	controllers "github.com/Juwon-Yun/goRest/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Resister)
}
