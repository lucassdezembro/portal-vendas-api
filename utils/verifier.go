package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CompareHashes(hash1 string, hash2 string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash1), []byte(hash2))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
