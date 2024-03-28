package services

import (
	"github.com/lucassdezembro/portal-vendas-api/entities"
	models "github.com/lucassdezembro/portal-vendas-api/models/requests"
	"github.com/lucassdezembro/portal-vendas-api/repositories"
	"github.com/lucassdezembro/portal-vendas-api/utils"
	errors_utils "github.com/lucassdezembro/portal-vendas-api/utils/errors"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*entities.UserEntity, errors_utils.Error) {

	userList, err := s.GetUserByDocument(req.User.Document, false)
	if err != nil {
		return nil, err
	}

	if userList != nil {
		return nil, errors_utils.NewBadRequestError("user already exists")
	}

	hashedPassword, err := utils.HashAndSalt(req.User.Password)
	if err != nil {
		return nil, err
	}

	entity := entities.UserEntity{
		Name:     req.User.Name,
		Document: req.User.Document,
		Email:    req.User.Email,
		Phone:    req.User.Phone,
		Password: hashedPassword,
	}

	entity.GenerateId()

	err = s.repository.CreateUser(entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (s *UserService) GetUser(id string) (*entities.UserEntity, error) {
	return s.repository.GetUser(id)
}

func (s *UserService) GetAllUsers() ([]entities.UserEntity, error) {
	return s.repository.QueryUsers(models.QueryUsersRequest{})
}

func (s *UserService) GetUserByDocument(document string, returnError bool) (*entities.UserEntity, error) {
	users, err := s.repository.QueryUsers(models.QueryUsersRequest{
		Document: document,
	})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		if returnError {
			return nil, errors_utils.NewResourceNotFoundError("user not found")
		}
		return nil, nil
	}

	return &users[0], nil
}

func (s *UserService) GetUserByEmail(email string, returnError bool) (*entities.UserEntity, error) {
	users, err := s.repository.QueryUsers(models.QueryUsersRequest{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		if returnError {
			return nil, errors_utils.NewResourceNotFoundError("user not found")
		}
		return nil, nil
	}

	return &users[0], nil
}
