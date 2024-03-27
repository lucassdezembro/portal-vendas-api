package services

import (
	"github.com/lucassdezembro/portal-vendas-api/entities"
	models "github.com/lucassdezembro/portal-vendas-api/models/requests"
	"github.com/lucassdezembro/portal-vendas-api/repositories"
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

	userList, err := s.repository.QueryUsers(models.QueryUsersRequest{
		Document: req.User.Document,
	})
	if err != nil {
		return nil, err
	}

	if len(userList) > 0 {
		return nil, errors_utils.NewBadRequestError("user already exists")
	}

	entity := entities.UserEntity{
		Name:     req.User.Name,
		Document: req.User.Document,
		Email:    req.User.Email,
		Phone:    req.User.Phone,
		Password: req.User.Password,
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
