package handler

import (
	"backend-mono/cmd/middleware"
	"backend-mono/cmd/svc"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	router.GET("/greet/:user_id", GreetHandler(serverCtx))
	router.POST("/user", CreateUserHandler(serverCtx))
	router.GET("/user/:user_id", middleware.JWTMiddleware(serverCtx), GetUserByIDHandler(serverCtx))
}
