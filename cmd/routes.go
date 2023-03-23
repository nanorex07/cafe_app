package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanorex07/cafe_app/routers"
)

func SetupRoutes(app *fiber.App) {

	router := app.Group("/auth")
	routers.SetupAuth(router)

	router = app.Group("/menu")
	routers.SetupMenu(router)
}
