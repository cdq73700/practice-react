package test

import (
	"github.com/gofiber/fiber/v2"
)

func GetTest(c *fiber.Ctx) error {
	return c.SendString("API TEST!!!")
}
