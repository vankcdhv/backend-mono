package svc

import (
	"backend-mono/cmd/config"
	"backend-mono/cmd/database/mysql"
	"backend-mono/cmd/database/repo"
)

type ServiceContext struct {
	Config     config.Config
	UserRepo   repo.UserRepo
	JWTService JWTService
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRepo, err := mysql.NewUserDB(c)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		UserRepo:   userRepo,
		JWTService: NewJWTService(),
	}
}
