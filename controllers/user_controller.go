package controllers

import (
	"github.com/gofiber/fiber"
	models "github.com/lucassdezembro/portal-vendas-api/models/requests"
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

	req := &struct {
		User struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	if err := c.BodyParser(req); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}

	serviceReq := models.CreateUserRequest{
		User: models.CreateUserRequest_User{
			Name:     req.User.Name,
			Email:    req.User.Email,
			Phone:    req.User.Phone,
			Password: req.User.Password,
		},
	}

	result, err := u.userService.CreateUser(serviceReq)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": *result,
	})
}

func (u *UserController) UpdateUser(c *fiber.Ctx) {
	c.SendString("Update user")
}
