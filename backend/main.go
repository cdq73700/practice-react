package main

import (
	v1 "backend/api/v1"
	"backend/database"
	"backend/db/migration"
	"backend/internal/user"
	model "backend/model"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

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

	app := fiber.New()

	// Default config
	app.Use(cors.New())

	userRepository := user.NewUserRepository(model.UserCollectionClient)
	userService := user.NewUserService(userRepository)
	user.NewUserHandler(app.Group("/api/v1/users"), userService)


	data,_ := database.GetMongoDbCollection("test","users")

	v1.NewUserHandler(app.Group("/test"), data.Collection)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Listen(":3000")
}
