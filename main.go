package main

import (
	dataBase "github.com/Juwon-Yun/goRest/database"
	"github.com/Juwon-Yun/goRest/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	dataBase.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":4000")
}
