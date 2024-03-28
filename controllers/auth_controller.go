package controllers

import (
	"github.com/gofiber/fiber"
	models "github.com/lucassdezembro/portal-vendas-api/models/requests"
	"github.com/lucassdezembro/portal-vendas-api/services"
	"github.com/lucassdezembro/portal-vendas-api/utils"
)

type AuthController struct {
	userService *services.UserService
	authService *services.AuthService
}

func NewAuthController(userService *services.UserService, authService *services.AuthService) *AuthController {
	return &AuthController{
		userService: userService,
		authService: authService,
	}
}

func (a *AuthController) Register(c *fiber.Ctx) {

	req := &struct {
		User struct {
			Name     string `json:"name"`
			Document string `json:"document"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	if err := c.BodyParser(req); err != nil {
		utils.HandleErrorData(c, err, fiber.StatusBadRequest)
		return
	}

	serviceReq := models.CreateUserRequest{
		User: models.CreateUserRequest_User{
			Name:     req.User.Name,
			Document: req.User.Document,
			Email:    req.User.Email,
			Phone:    req.User.Phone,
			Password: req.User.Password,
		},
	}

	result, err := a.userService.CreateUser(serviceReq)
	if err != nil {
		utils.HandleErrorData(c, err, 0)
		return
	}

	token, err := a.authService.RegisterUser(*result)
	if err != nil {
		utils.HandleErrorData(c, err, 0)
		return
	}

	response := map[string]interface{}{
		"user":  result,
		"token": token,
	}

	utils.HandleSuccessData(c, response, fiber.StatusCreated)
}
