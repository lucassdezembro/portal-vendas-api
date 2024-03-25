package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/routes"
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

	routes.SetupRoutes(app)

	fmt.Printf("Server is running on port: %d...", PORT)

	app.Listen(fmt.Sprintf(":%d", PORT))
}
