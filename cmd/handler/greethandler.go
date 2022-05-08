package handler

import (
	"backend-mono/cmd/logic"
	"backend-mono/cmd/svc"
	"backend-mono/cmd/types"
	"backend-mono/core/http_request"
	"backend-mono/core/http_response"
	"backend-mono/core/logger"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GreetHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("greet-service"))
		// Init log helper with context (have trace_id)
		logHelper := logger.NewContextLog(ctx)
		// New object logic (all logic code will implement in this object)
		greetLogic := logic.NewGreetLogic(ctx, logHelper, svcCtx)

		// Parse request mysql from request
		request := &types.GreetRequest{}
		err := http_request.BindUri(c, request)
		if err != nil {
			logHelper.Errorw("Failed while parse greet request", "extra_readable_data", err.Error())
			http_response.ResponseJSON(c, http.StatusBadGateway, err.Error())
			return
		}

		// Call functions in logic to process
		resp, err := greetLogic.GreetMessage(request)

		// Response mysql to client
		// Include http status code (resp.Code) and mysql resp (resp)
		// If have error when processing logic, err will not nil and resp have error message
		http_response.ResponseJSON(c, resp.Code, resp)
	}
}
