package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/controllers"
	"github.com/lucassdezembro/portal-vendas-api/db"
	"github.com/lucassdezembro/portal-vendas-api/middlewares"
	"github.com/lucassdezembro/portal-vendas-api/repositories"
	"github.com/lucassdezembro/portal-vendas-api/routes"
	"github.com/lucassdezembro/portal-vendas-api/services"
)

var (
	PORT int
)

func init() {
	flag.IntVar(&PORT, "port", 8080, "api service port")
	flag.Parse()
}

func main() {

	app := fiber.New()

	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	//setup repositories
	userRepository := repositories.NewUserRepository(dbConnection)
	authRepository := repositories.NewAuthRepository()

	//setup services
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(authRepository)

	//setup middlewares
	tokenMiddleware := middlewares.NewTokenMiddleware(authService, userService)

	routesOptions := map[string]interface{}{
		"user": map[string]interface{}{
			"controller": controllers.NewUserController(userService),
		},
		"auth": map[string]interface{}{
			"controller": controllers.NewAuthController(userService, authService),
		},
		"advertisement": map[string]interface{}{
			"middlewares": map[string]interface{}{
				"tokenMiddleware": tokenMiddleware,
			},
		},
	}

	routes.SetupRoutes(app, routesOptions)

	fmt.Printf("Server is running on port: %d...", PORT)

	app.Listen(fmt.Sprintf(":%d", PORT))
}
