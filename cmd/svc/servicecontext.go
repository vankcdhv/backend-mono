package svc

import (
	"backend-mono/cmd/config"
	"backend-mono/database/repo"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config, userRepo repo.UserRepo) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRepo: userRepo,
	}
}
