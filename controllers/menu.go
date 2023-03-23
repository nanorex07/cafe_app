package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nanorex07/cafe_app/database"
	"github.com/nanorex07/cafe_app/models"
)

func MenuCreate(c *fiber.Ctx) error {
	uid, _ := strconv.Atoi(c.Locals("user").(string))

	menuReq := &models.MenuCreateReq{}

	err := c.BodyParser(menuReq)
	if err != nil {
		return err
	}
	menu := &models.Menu{
		Name:        menuReq.Name,
		Description: menuReq.Description,
		UserID:      uint(uid),
	}
	database.DB.Db.Create(menu)

	return c.JSON(fiber.Map{
		"message": "created",
	})
}

func GetMenus(c *fiber.Ctx) error {
	uid, _ := strconv.Atoi(c.Locals("user").(string))

	var ret []models.Menu

	database.DB.Db.Preload("Items").Find(&ret, "user_id = ?", uint(uid))
	return c.JSON(ret)
}

func ItemCreate(c *fiber.Ctx) error {
	menu_id, _ := strconv.Atoi(c.Params("menu_id"))
	itemReq := &models.ItemCreateRequest{}

	err := c.BodyParser(itemReq)
	fmt.Println(itemReq)
	if err != nil {
		return err
	}
	errors := models.ValidateStruct(itemReq)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	itemdb := models.Item{
		MenuID:      uint(menu_id),
		Name:        itemReq.Name,
		Description: itemReq.Description,
		Price:       itemReq.Price,
	}

	database.DB.Db.Create(&itemdb)

	return c.JSON(fiber.Map{
		"message": "created",
	})
}

func GetItems(c *fiber.Ctx) error {
	menu_id, _ := strconv.Atoi(c.Params("menu_id"))

	var ret []models.Item
	database.DB.Db.Find(&ret, "menu_id = ?", uint(menu_id))
	return c.JSON(ret)
}

func AddImageToItem(c *fiber.Ctx) error {
	uid, _ := strconv.Atoi(c.Locals("user").(string))
	menu_id, _ := strconv.Atoi(c.Params("menu_id"))
	item_id, _ := strconv.Atoi(c.Params("item_id"))

	var item models.Item
	var menu models.Menu
	database.DB.Db.First(&item, "id = ? AND menu_id = ?", item_id, menu_id)
	database.DB.Db.First(&menu, "id = ? AND user_id = ?", menu_id, uid)

	if item.MenuID == 0 || menu.ID == 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "unauthorised or resource unavailable",
		})
	}
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	fileName := strconv.Itoa(item_id) + "_" + file.Filename
	c.SaveFile(file, "./images/"+fileName)
	item.ImageURL = fileName
	database.DB.Db.Updates(item)
	return c.JSON(item)
}
