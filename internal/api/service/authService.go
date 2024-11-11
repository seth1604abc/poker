package service

import (
	"poker/config"
	"poker/internal/api/dto"
	"poker/internal/api/repository"
	"poker/internal/errors"
	"poker/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserUid	string	`json:"userUid"`
	jwt.RegisteredClaims
}

type authService struct {
	userRepo repository.UserRepository
}

type AuthService interface {
	Login(data dto.LoginParams) (string, error)
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{ userRepo: userRepo }
}

func (s *authService) Login(data dto.LoginParams) (string, error) {
	user, userErr := s.userRepo.FindByAccount(data.Account)
	if userErr != nil {
		return "", errors.NewDatabaseError("Login")
	}

	if user == nil {
		return "", errors.NewNotFoundError("Login")
	}

	valid, validErr := util.CheckPassword(data.Password, user.Password)
	if validErr != nil {
		return "", errors.NewInternalServerError("Login")
	}
	if !valid {
		return "", errors.NewPasswordValidError("Login")
	}
	userUid := strconv.Itoa(int(user.ID))

	token, err := s.GenerateAccessToken(userUid)
	if err != nil {
		return "", err
	}
	
	return token, nil
}

func (s *authService) GenerateAccessToken(userUid string) (string, error) {
	expiredTime := time.Now().Add(time.Duration(config.AppConfig.JWT.ExpiredTime) * time.Second)

	claims := &Claims{
		UserUid: userUid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.AppConfig.JWT.SecretKey)
}
