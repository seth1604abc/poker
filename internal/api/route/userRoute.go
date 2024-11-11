package route

import (
	"net/http"
	"poker/internal/api/controller"
	"poker/internal/api/dto"

	"github.com/gin-gonic/gin"
)

type userRoute struct{
	userController controller.UserController
}

type UserRoute interface{
	SetUp(engine *gin.Engine)
}

func NewUserRoute(userController controller.UserController) UserRoute {
	return &userRoute{
		userController: userController,
	}
}

func (r *userRoute) SetUp(engine *gin.Engine) {
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/", r.RegisterUser)
	}
}

// ============= route ==============

func (r *userRoute) RegisterUser(ctx *gin.Context) {
	var body dto.RegisterUserParams

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r.userController.RegisterUser(ctx, body)
}

