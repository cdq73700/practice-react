package main

import (
	"backend/api/v1/test"
	"backend/db"
	"backend/db/migration"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1/test", test.GetTest)
}

func main() {
	arg1 := os.Args[1]
	client, err := db.GetMongoDbConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	if arg1 == "migration" {

		migration.Migration(client)

		fmt.Println("run migration")
		return
	}

	if arg1 == "delete" {
		migration.CollectionDrop(client)

		fmt.Println("run detele")
		return
	}

	if arg1 == "drop" {
		migration.DatabaseDrop(client)

		fmt.Println("run drop")
		return
	}
	app := fiber.New()

	//client, err := db.GetMongoDbConnection()

	setUpRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
