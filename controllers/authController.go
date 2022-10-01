package authController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/vannguyen2606/poseidon-core/database"
	"github.com/vannguyen2606/poseidon-core/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user models.User
	myEnv, _ := godotenv.Read()

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

	claims := jwt.MapClaims{
		"password": data["password"],
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, errToken := token.SignedString([]byte(myEnv["SECRET_KEY"]))

	if errToken != nil {
		return c.SendString("gen jwt failed!")
	}

	return c.JSON(fiber.Map{
		"token": ss,
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
