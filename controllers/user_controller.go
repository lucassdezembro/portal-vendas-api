package controllers

import "github.com/gofiber/fiber"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
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
