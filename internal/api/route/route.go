package route

import "github.com/gin-gonic/gin"

type Route struct {
	userRouter UserRoute
}

func NewRoute(userRouter UserRoute) *Route {
	return &Route{
		userRouter: userRouter,
	}
}

func (r *Route) SetUp(engine *gin.Engine) {
	r.userRouter.SetUp(engine)
}