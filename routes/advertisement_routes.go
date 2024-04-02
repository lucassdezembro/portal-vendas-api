package routes

import (
	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/middlewares"
)

func AdvertisementRoutes(app *fiber.App, advertisementOptions map[string]interface{}) {

	tokenMiddleware := advertisementOptions["middlewares"].(map[string]interface{})["tokenMiddleware"].(*middlewares.TokenMiddleware)

	app.Use("/advertisement", tokenMiddleware.VerifyToken)

	app.Post("/advertisement", func(c *fiber.Ctx) {
		c.Send("Create advertisement")
	})
}
