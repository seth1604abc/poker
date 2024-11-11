package controller

import (
	"net/http"
	"poker/internal/api/dto"
	"poker/internal/api/service"
	"poker/internal/errors"
	"poker/util"
	"poker/util/message"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(ctx *gin.Context, body dto.RegisterUserParams)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

func (c *userController) RegisterUser(ctx *gin.Context, body dto.RegisterUserParams) {
	// 驗證
	validAccount := util.IsValidAccount(body.Account)
	if !validAccount {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message.NotValidError("account"),
		})
		return
	}

	validPassword := util.IsValidPassword(body.Password)
	if !validPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message.NotValidError("password"),
		})
		return
	}

	validEmail := util.IsValidEmail(body.Email)
	if !validEmail {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": message.NotValidError("email"),
		})
		return
	}

	registerErr := c.userService.RegisterUser(dto.CreateUserParams{
		Account: body.Account,
		Email: body.Email,
		Password: body.Password,
		Name: body.Name,
	})

	if registerErr != nil {
		if _, ok := registerErr.(*errors.DatabaseError); ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": message.InternalServerError(),
			})
			return
		} else if _, ok := registerErr.(*errors.DuplicateAccountError); ok {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": message.DuplicateAccountError(),
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": message.InternalServerError(),
			})
			return
		}
	}
}