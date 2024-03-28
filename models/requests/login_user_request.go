package models

import (
	"github.com/lucassdezembro/portal-vendas-api/entities"
)

type LoginUserRequest struct {
	User     entities.UserEntity `json:"user"`
	Password string              `json:"password"`
}
