package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/lucassdezembro/portal-vendas-api/entities"
	models "github.com/lucassdezembro/portal-vendas-api/models/requests"
	"github.com/lucassdezembro/portal-vendas-api/services"
	"github.com/lucassdezembro/portal-vendas-api/utils"
	errors_utils "github.com/lucassdezembro/portal-vendas-api/utils/errors"
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

func (a *AuthController) Login(c *fiber.Ctx) {

	req := &struct {
		User struct {
			Document string `json:"document"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}{}

	if err := c.BodyParser(req); err != nil {
		utils.HandleErrorData(c, err, fiber.StatusBadRequest)
		return
	}

	if req.User.Password == "" {
		utils.HandleErrorData(c, errors_utils.NewBadRequestError("missing user.password"))
		return
	}

	var (
		user *entities.UserEntity
		err  error
	)

	if req.User.Document != "" {
		user, err = a.userService.GetUserByDocument(req.User.Document, true)
	} else if req.User.Email != "" {
		user, err = a.userService.GetUserByEmail(req.User.Email, true)
	} else {
		utils.HandleErrorData(c, errors_utils.NewBadRequestError("missing user.document or user.email"))
		return
	}

	// nota: os erros da busca de usuário e login são genéricos e pouco informativos, para evitar dar qualquer tipo de informação para um possível invasor.

	if err != nil {
		utils.HandleErrorData(c, errors_utils.NewBadRequestError("invalid credentials"))
		return
	}

	serviceReq := models.LoginUserRequest{
		Password: req.User.Password,
		User:     *user,
	}

	token, err := a.authService.Login(serviceReq)
	if err != nil {
		utils.HandleErrorData(c, errors_utils.NewBadRequestError("invalid credentials"))
		return
	}

	response := map[string]interface{}{
		"user":  *user,
		"token": token,
	}

	utils.HandleSuccessData(c, response, fiber.StatusCreated)
}
