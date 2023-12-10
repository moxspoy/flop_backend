package user_detail_router

import (
	user_controller2 "flop/controllers/api/v1/user_detail_controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterUserDetailRouter(v1Router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	userRouter := v1Router.Group("/user-detail")
	userRouter.Use(authMiddleware.MiddlewareFunc())
	{
		userRouter.GET("/info", user_controller2.GetUserDetail)
	}
}