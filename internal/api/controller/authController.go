package controller

import (
	"poker/internal/api/dto"
	"poker/internal/api/service"
	"poker/internal/errors"

	"github.com/gin-gonic/gin"
)

type authController struct{
	authService service.AuthService
}

type AuthController interface{}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{authService: authService}
}

func (c *authController) Login(ctx *gin.Context, body dto.LoginParams) {
	token, err := c.authService.Login(body)

	if err != nil {
		if _, ok := err.(*errors.DatabaseError); ok {
			
		}
	}
}