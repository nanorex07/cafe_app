package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/nanorex07/cafe_app/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ping": "pong",
		})
	})

	SetupRoutes(app)

	app.Static("/", "./public")

	log.Fatal(app.Listen(":8000"))
}
