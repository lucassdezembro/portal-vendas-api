package services

import "github.com/lucassdezembro/portal-vendas-api/repositories"

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}
