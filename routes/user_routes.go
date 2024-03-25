package routes

import "github.com/gofiber/fiber"

func UserRoutes(app *fiber.App) {

	app.Get("/users", func(c *fiber.Ctx) {
		c.SendString("All users")
	})

	app.Get("/users/:id", func(c *fiber.Ctx) {
		c.SendString("Get user by id")
	})

	app.Post("/users", func(c *fiber.Ctx) {
		c.SendString("Create user")
	})

	app.Put("/users/:id", func(c *fiber.Ctx) {
		c.SendString("Update user")
	})
}
