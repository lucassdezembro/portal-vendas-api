package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber"
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

	app.Get("/", func(c *fiber.Ctx) {
		c.SendString("Hello, World!")
	})

	fmt.Printf("Server is running on port: %d...", PORT)

	app.Listen(fmt.Sprintf(":%d", PORT))
}
