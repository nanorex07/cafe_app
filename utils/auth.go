package utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nanorex07/cafe_app/database"
	"github.com/nanorex07/cafe_app/models"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenCookie := c.Cookies("auth")

	token, err := jwt.ParseWithClaims(tokenCookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	},
	)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	c.Locals("user", claims.Issuer)
	return c.Next()
}

func GetUser(id int) models.User {
	var user models.User
	database.DB.Db.Where("id = ?", id).First(&user)
	return user
}
