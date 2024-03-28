package services

import (
	"github.com/lucassdezembro/portal-vendas-api/entities"
	models "github.com/lucassdezembro/portal-vendas-api/models/requests"
	"github.com/lucassdezembro/portal-vendas-api/repositories"
	"github.com/lucassdezembro/portal-vendas-api/utils"
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

func (s *AuthService) Login(loginReq models.LoginUserRequest) (string, error) {

	err := utils.CompareHashes(loginReq.User.Password, loginReq.Password)
	if err != nil {
		return "", err
	}

	token, err := s.AuthRepository.GenerateJWT(loginReq.User)
	if err != nil {
		return "", err
	}

	return token, nil

}
