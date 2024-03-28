package routes

import (
	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/controllers"
)

func UserRoutes(app *fiber.App, userOptions map[string]interface{}) {

	controller := userOptions["controller"].(*controllers.UserController)

	app.Get("/users", controller.GetAllUsers)

	app.Get("/users/:id", controller.GetUserById)

	app.Put("/users/:id", controller.UpdateUser)
}
