package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanorex07/cafe_app/controllers"
)

func SetupAuth(app fiber.Router) {

	app.Post("/register", controllers.AuthRegister)

	app.Post("/login", controllers.AuthLogin)

	app.Get("/logout", controllers.AuthLogout)

}
