package controllers

import (
	"github.com/Juwon-Yun/goRest/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Resister(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// hased password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := model.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	return c.JSON(user)
}
