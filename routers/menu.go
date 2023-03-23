package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nanorex07/cafe_app/controllers"
	"github.com/nanorex07/cafe_app/utils"
)

func SetupMenu(app fiber.Router) {
	// get and create menus
	app.Post("/", utils.AuthMiddleware, controllers.MenuCreate)
	app.Get("/", utils.AuthMiddleware, controllers.GetMenus)

	// get and create items
	app.Post("/:menu_id/item", utils.AuthMiddleware, controllers.ItemCreate)
	app.Get("/:menu_id/item", utils.AuthMiddleware, controllers.GetItems)

	// upload image for an item
	app.Post("/:menu_id/item/:item_id/upload_image", utils.AuthMiddleware, controllers.AddImageToItem)
}
