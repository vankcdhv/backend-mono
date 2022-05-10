package svc

import (
	"backend-mono/cmd/config"
	"backend-mono/cmd/database/repo"
	"backend-mono/cmd/service"
)

type ServiceContext struct {
	Config     config.Config
	UserRepo   repo.UserRepo
	JWTService service.JWTService
}

func NewServiceContext(c config.Config, userRepo repo.UserRepo, jwtService service.JWTService) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRepo:   userRepo,
		JWTService: jwtService,
	}
}
