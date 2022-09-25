package authController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vannguyen2606/poseidon-core/database"
	"github.com/vannguyen2606/poseidon-core/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user models.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "login successfully!",
	})
}

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
