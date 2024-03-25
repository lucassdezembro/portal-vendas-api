package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/controllers"
	"github.com/lucassdezembro/portal-vendas-api/db"
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

	//setup services
	userService := services.NewUserService(userRepository)

	routesOptions := map[string]interface{}{
		"user": map[string]interface{}{
			"controller": controllers.NewUserController(userService),
		},
	}

	routes.SetupRoutes(app, routesOptions)

	fmt.Printf("Server is running on port: %d...", PORT)

	app.Listen(fmt.Sprintf(":%d", PORT))
}
