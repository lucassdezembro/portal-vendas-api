package services

import (
	"github.com/lucassdezembro/portal-vendas-api/entities"
	"github.com/lucassdezembro/portal-vendas-api/repositories"
)

type AuthService struct {
	AuthRepository *repositories.AuthRepository
}

func NewAuthService(authRepository *repositories.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}

func (s *AuthService) RegisterUser(user entities.UserEntity) (string, error) {

	token, err := s.AuthRepository.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Login() {

}
