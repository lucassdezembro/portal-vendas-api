package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashes(hash string, plain string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		return err
	}

	return nil
}
