package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nanorex07/cafe_app/database"
	"github.com/nanorex07/cafe_app/models"
	"golang.org/x/crypto/bcrypt"
)

func AuthRegister(c *fiber.Ctx) error {

	user := &models.User{}

	err := c.BodyParser(user)
	if err != nil {
		return err
	}
	errors := models.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(pass)

	database.DB.Db.Create(user)

	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "User already exists",
		})
	}

	return c.JSON(user)

}

func AuthLogin(c *fiber.Ctx) error {
	user := &models.UserLogin{}

	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	var userdb models.User

	database.DB.Db.Where("email = ?", user.Email).First(&userdb)

	if userdb.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": "User not found",
			},
		)
	}
	fmt.Println(user.Password)
	err = bcrypt.CompareHashAndPassword([]byte(userdb.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "incorrect password" + err.Error(),
			},
		)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userdb.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Can't login " + err.Error(),
		})
	}
	cookie := fiber.Cookie{
		Name:    "auth",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}
	c.Cookie(&cookie)

	return c.JSON(userdb)
}

func AuthLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "logout success",
	})
}
