package repository

import (
	"errors"
	"poker/internal/api/dto"
	"poker/internal/api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user dto.CreateUserParams) error
	FindByAccount(account string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db:db}
}

func (r *userRepository) CreateUser(user dto.CreateUserParams) error {
	return r.db.Create(&model.User{
		Name: user.Name,
		Account: user.Account,
		Password: user.Password,
		Email: user.Email,
	}).Error
}

func (r *userRepository) FindByAccount(account string) (*model.User, error) {
	var user model.User

	result := r.db.Where(&model.User{ Account: account }).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}	
	}
	return &user, nil
}
