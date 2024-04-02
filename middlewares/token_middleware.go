package middlewares

import (
	"strings"

	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/services"
)

type TokenMiddleware struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewTokenMiddleware(authService *services.AuthService, userService *services.UserService) *TokenMiddleware {
	return &TokenMiddleware{
		authService: authService,
		userService: userService,
	}
}

func (m *TokenMiddleware) VerifyToken(c *fiber.Ctx) {

	token := c.Get("Authorization")

	if token == "" {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return
	}

	if strings.Contains(token, "Bearer") {
		token = strings.Replace(token, "Bearer ", "", -1)
	}

	isTokenValid, err := m.authService.VerifyToken(token)
	if err != nil || !isTokenValid {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return
	}

	decodedToken, err := m.authService.DecodeToken(token)
	if err != nil {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return
	}

	userDocument := ""

	for key, value := range decodedToken {
		if key == "userDocument" {
			userDocument = value.(string)
			break
		}
	}

	user, err := m.userService.GetUserByDocument(userDocument, true)
	if err != nil {
		c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return
	}

	c.Locals("currentUser", user)

	c.Next()
}
