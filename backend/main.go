package main

import (
	"backend/api/v1/test"
	"backend/api/v1/users"
	"backend/db/migration"
	model "backend/model"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func SetUpRoutes() {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	usersApi := v1.Group("/users")

	v1.Get("/test", test.GetTest)

	usersApi.Get("/", func(c *fiber.Ctx) error { return users.UserList(c) })
	usersApi.Get("/:id?", func(c *fiber.Ctx) error { return users.UserFind(c) })
	
	usersApi.Post("/add", func(c *fiber.Ctx) error { return users.UserAdd(c) })
	usersApi.Put("/update", func(c *fiber.Ctx) error { return users.UserUpdate(c) })
	usersApi.Delete("/delete", func(c *fiber.Ctx) error { return users.UserDelete(c) })

}

func RunCommand (arg1 string) string {
	var message string
	switch(arg1) {
	case "migration":
		migration.Migration()
		message = "run migration"
		break
	case "delete":
		migration.CollectionDelete()
		message ="run delete"
		break
	case "drop":
		migration.DatabaseDrop()
		message = "run drop"
		break
	default:
		message = ""
		break
	}

	return message
}

func main() {
	var arg1 = ""
	if len(os.Args) == 2 {
		arg1 = os.Args[1]
	}

	model.Initialized()

	msg := RunCommand(arg1)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	app = fiber.New()

	// Default config
	app.Use(cors.New())

	SetUpRoutes()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Listen(":3000")
}
