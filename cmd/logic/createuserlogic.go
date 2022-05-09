package logic

import (
	"backend-mono/cmd/svc"
	"backend-mono/cmd/types"
	"backend-mono/database/model"
	"backend-mono/database/repo"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
	"strconv"
)

type CreateUserLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	logHelper *log.Helper
}

func NewCreateUserLogic(ctx context.Context, logHelper *log.Helper, svcCtx *svc.ServiceContext) CreateUserLogic {
	return CreateUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		logHelper: logHelper,
	}
}

func (l *CreateUserLogic) CreateUser(input *types.CreateUserRequest) (*types.CreateUserResponse, error) {
	l.logHelper.Infof("Start process greet message, input: %+v", input)
	data := &repo.CreateUserIn{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	resp, err := l.svcCtx.UserRepo.Create(l.ctx, data)
	if err != nil {
		return &types.CreateUserResponse{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		}, err
	}

	return &types.CreateUserResponse{
		Code: http.StatusOK,
		Data: &model.User{
			ID:        strconv.FormatInt(resp.ID, 10),
			FirstName: data.FirstName,
			LastName:  data.LastName,
		},
	}, nil
}
