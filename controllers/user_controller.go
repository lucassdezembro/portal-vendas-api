package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/services"
	"github.com/lucassdezembro/portal-vendas-api/utils"
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

	result, err := u.userService.GetAllUsers()
	if err != nil {
		utils.HandleErrorData(c, err)
		return
	}

	utils.HandleSuccessData(c, result, fiber.StatusOK)
}

func (u *UserController) GetUserById(c *fiber.Ctx) {

	id := c.Params("id")

	result, err := u.userService.GetUser(id)
	if err != nil {
		utils.HandleErrorData(c, err, fiber.StatusInternalServerError)
		return
	}

	utils.HandleSuccessData(c, result, fiber.StatusOK)
}

func (u *UserController) UpdateUser(c *fiber.Ctx) {
	c.SendString("Update user")
}
