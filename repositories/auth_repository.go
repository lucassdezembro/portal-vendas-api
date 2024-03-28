package repositories

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucassdezembro/portal-vendas-api/entities"
)

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) GenerateJWT(user entities.UserEntity) (string, error) {

	jwtExpiresIn, err := strconv.Atoi(os.Getenv("JWT_MINUTES_EXPIRES_IN"))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userDocument": user.Document,
			"exp":          time.Now().Add(time.Minute * time.Duration(jwtExpiresIn)).Unix(),
		})

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
