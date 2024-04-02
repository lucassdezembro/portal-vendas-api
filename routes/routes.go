package routes

import "github.com/gofiber/fiber"

func SetupRoutes(app *fiber.App, options map[string]any) {
	UserRoutes(app, options["user"].(map[string]interface{}))
	AuthRoutes(app, options["auth"].(map[string]interface{}))
	AdvertisementRoutes(app, options["advertisement"].(map[string]interface{}))
}
