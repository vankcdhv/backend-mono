package logic

import (
	"backend-mono/cmd/svc"
	"backend-mono/cmd/types"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type GetUserByIDLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewGetUserByIDLogic(ctx context.Context, logHelper *log.Helper, svcCtx *svc.ServiceContext) GetUserByIDLogic {
	return GetUserByIDLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *GetUserByIDLogic) GetUserByID(input *types.GetUserByIDRequest) (*types.GetUserByIDResponse, error) {
	l.logHelper.Infof("Start process greet message, input: %+v", input)
	user, err := l.svcCtx.UserRepo.FindByID(l.ctx, input.UserID)
	if err != nil {
		l.logHelper.Errorf("Failed while get user by id, error: %s", err.Error())
		return &types.GetUserByIDResponse{
			Code:  http.StatusBadRequest,
			Data:  nil,
			Error: err.Error(),
		}, err
	}
	return &types.GetUserByIDResponse{
		Code: http.StatusOK,
		Data: &types.UserData{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Error: "",
	}, nil
}
