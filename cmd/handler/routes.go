package handler

import (
	"backend-mono/cmd/svc"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	router.GET("/greet/:user_id", GreetHandler(serverCtx))
}
