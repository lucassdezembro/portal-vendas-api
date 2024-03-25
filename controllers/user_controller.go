package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetAllUsers(c *fiber.Ctx) {
	c.SendString("All users")
}

func (u *UserController) GetUserById(c *fiber.Ctx) {
	c.SendString("Get user by id")
}

func (u *UserController) CreateUser(c *fiber.Ctx) {
	c.SendString("Create user")
}

func (u *UserController) UpdateUser(c *fiber.Ctx) {
	c.SendString("Update user")
}
