package logger

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

func NewContextLog(ctx context.Context) *log.Helper {
	logHelper := log.NewHelper(log.DefaultLogger)
	logHelper.WithContext(ctx)
	return logHelper
}

func GenerateTraceID(serviceName string) string {
	now := time.Now()
	traceID := fmt.Sprintf("%s-%d", serviceName, now.UnixNano())
	return traceID
}
