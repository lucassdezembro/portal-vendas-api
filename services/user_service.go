package services

import (
	"github.com/lucassdezembro/portal-vendas-api/entities"
	models "github.com/lucassdezembro/portal-vendas-api/models/requests"
	"github.com/lucassdezembro/portal-vendas-api/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*entities.UserEntity, error) {

	entity := entities.UserEntity{
		Name:     req.User.Name,
		Email:    req.User.Email,
		Phone:    req.User.Phone,
		Password: req.User.Password,
	}

	entity.GenerateId()

	err := s.repository.CreateUser(entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
