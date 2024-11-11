package route

import (
	"net/http"
	"poker/internal/api/controller"
	"poker/internal/api/dto"

	"github.com/gin-gonic/gin"
)

type authRoute struct {
	authController controller.AuthController
}

type AuthRoute interface{}

func NewAuthRoute(authController controller.AuthController) AuthRoute {
	return &authRoute{authController: authController}
}

// =============== route ==============

func (r *authRoute) Login(ctx *gin.Context) {
	var body dto.LoginParams

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.authController
}