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

func (r *UserRepository) GetUser(id string) (*entities.UserEntity, error) {

	user := &entities.UserEntity{}

	tx := r.db.Where("id = ?", id).First(user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}
