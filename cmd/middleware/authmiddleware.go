package middleware

import (
	"backend-mono/cmd/svc"
	"backend-mono/core/logger"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

const BearerSchema = "Bearer "

func JWTMiddleware(svcContext *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "trace_id", logger.GenerateTraceID("auth-middleware"))
		ctxLogger := logger.NewContextLog(ctx)
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := svcContext.JWTService.ValidateToken(tokenString)
		if err != nil {
			ctxLogger.Errorf("Failed while validate authentication token %s", err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			userID, ok := claims["user_id"]
			if !ok {
				ctxLogger.Infof("User id invalid")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Request.Header.Add("x-user-id", fmt.Sprintf("%d", userID))
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
