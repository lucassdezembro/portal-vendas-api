package routes

import (
	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/controllers"
)

func AuthRoutes(app *fiber.App, authOptions map[string]interface{}) {

	controller := authOptions["controller"].(*controllers.AuthController)

	app.Post("/auth/register", controller.Register)
	app.Post("/auth/login", controller.Login)
}
