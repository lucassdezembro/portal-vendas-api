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

func (r *AuthRepository) VerifyJWT(tokenString string) (bool, error) {

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func (r *AuthRepository) DecodeJWT(tokenString string) (jwt.MapClaims, error) {

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}
