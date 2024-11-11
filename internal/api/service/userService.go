package service

import (
	"poker/internal/api/dto"
	"poker/internal/api/repository"
	"poker/internal/errors"
	"poker/util"
)

type UserService interface{
	RegisterUser (data dto.CreateUserParams) error
}

type userService struct{
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo:userRepo}
}

func (s *userService) RegisterUser (data dto.CreateUserParams) error {
	// 檢查帳號重複
	user, userErr := s.userRepo.FindByAccount(data.Account)
	if userErr != nil {
		return errors.NewDatabaseError("RegisterUser")
	}
	if user != nil {
		return errors.NewDuplicateAccountError("RegisterUser")
	}

	hashedPassword, hashedPasswordError := util.HashPassword(data.Password)
	if hashedPasswordError != nil {
		return errors.NewInternalServerError("RegisterUser")
	}

	createUserError := s.userRepo.CreateUser(dto.CreateUserParams{
		Account: data.Account,
		Name: data.Name,
		Password: hashedPassword,
		Email: data.Email,
	})
	if createUserError != nil {
		return errors.NewDatabaseError("RegisterUser")
	}

	return nil
}