package repositories

import (
	"github.com/lucassdezembro/portal-vendas-api/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user entities.UserEntity) error {

	r.db.AutoMigrate(&entities.UserEntity{})

	tx := r.db.Create(user)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
