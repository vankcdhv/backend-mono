package logic

import (
	"backend-mono/cmd/svc"
	"backend-mono/cmd/types"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type GreetLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewGreetLogic(ctx context.Context, logHelper *log.Helper, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *GreetLogic) GreetMessage(input *types.GreetRequest) (*types.GreetResponse, error) {
	l.logHelper.Infof("Start process greet message, input: %+v", input)
	user, err := l.svcCtx.UserRepo.FindByID(l.ctx, input.UserID)
	if err != nil {
		return &types.GreetResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		}, err
	}
	message := fmt.Sprintf("%s%s", "Hello user ", user.LastName)
	return &types.GreetResponse{
		Code:    http.StatusOK,
		Message: message,
	}, nil
}
