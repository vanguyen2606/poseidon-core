package routes

import (
	"github.com/vannguyen2606/poseidon-core/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("auth/login", authController.Login)
	app.Post("auth/register", authController.Register)
}