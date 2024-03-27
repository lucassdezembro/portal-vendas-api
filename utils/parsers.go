package utils

import (
	errors_utils "github.com/lucassdezembro/portal-vendas-api/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(str string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", errors_utils.NewInternalServerError("error hashing password")
	}

	return string(hash), nil
}
