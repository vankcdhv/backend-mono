package handler

import (
	"backend-mono/cmd/middleware"
	"backend-mono/cmd/svc"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	router.GET("/greet/:user_id", middleware.JWTMiddleware(serverCtx), GreetHandler(serverCtx))
	router.POST("/user", CreateUserHandler(serverCtx))
	router.GET("/user/:user_id", GetUserByIDHandler(serverCtx))
}
