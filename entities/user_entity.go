package entities

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

type UserEntity struct {
	Id       string `json:"id" gorm:"type:varchar(255);primaryKey;not null;"`
	Document string `json:"document" gorm:"type:varchar(14);not null;unique"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Phone    string `json:"phone" gorm:"type:varchar(20);not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
}

func (entity *UserEntity) GenerateId() {

	base := entity.Document + time.Now().UTC().Format(time.RFC3339)

	hash := sha1.New()
	hash.Write([]byte(base))

	entity.Id = string(hex.EncodeToString(hash.Sum(nil)))
}
