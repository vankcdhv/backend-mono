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

func GetUserByIDHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add trace_id to context
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("get-user-by-id-api"))
		// Init log helper with context (have trace_id)
		logHelper := logger.NewContextLog(ctx)
		// New object logic (all logic code will implement in this object)
		getUserByIDLogic := logic.NewGetUserByIDLogic(ctx, logHelper, svcCtx)
		// Parse data from request
		request := &types.GetUserByIDRequest{}
		err := http_request.BindUri(c, request)
		if err != nil {
			logHelper.Errorw("Failed while parse greet request", "extra_readable_data", err.Error())
			http_response.ResponseJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		// Call functions in logic to process
		resp, err := getUserByIDLogic.GetUserByID(request)

		// Response mysql to client
		// Include http status code (resp.Code) and mysql resp (resp)
		// If have error when processing logic, err will not nil and resp have error message
		http_response.ResponseJSON(c, resp.Code, resp)
	}
}
