package repository

import (
	"poker/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db:db}
}

func (r *userRepository) createUser(user *model.User) error {
	return r.db.Create(user).Error
}
