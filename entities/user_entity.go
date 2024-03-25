package entities

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

type UserEntity struct {
	Id       string `gorm:"type:varchar(255);primaryKey;not null;"`
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Phone    string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
}

func (entity *UserEntity) GenerateId() {

	base := entity.Name + time.Now().UTC().Format(time.RFC3339)

	hash := sha1.New()
	hash.Write([]byte(base))

	entity.Id = string(hex.EncodeToString(hash.Sum(nil)))
}
