package main

import (
	"backend/api/v1/test"
	"backend/api/v1/users"
	"backend/db"
	"backend/db/migration"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

func setUpRoutes(app *fiber.App, db *mongo.Database) {
	app.Get("/api/v1/test", test.GetTest)
	app.Get("/api/v1/users/:id?", func(c *fiber.Ctx) error { return users.GetUsers(c, db) })
	
	app.Post("/api/v1/users/add", func(c *fiber.Ctx) error { return users.AddUser(c, db) })
}

func main() {
	var arg1 = ""
	if len(os.Args) == 2 {
		arg1 = os.Args[1]
	}

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
		migration.CollectionDelete(client)

		fmt.Println("run detele")
		return
	}

	if arg1 == "drop" {
		migration.DatabaseDrop(client)

		fmt.Println("run drop")
		return
	}
	app := fiber.New()

	// Default config
	app.Use(cors.New())

	setUpRoutes(app, client.Database("test"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Listen(":3000")
}
