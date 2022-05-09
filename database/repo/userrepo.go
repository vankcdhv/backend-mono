package repo

import (
	"backend-mono/database/model"
	"context"
)

type CreateUserIn struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateUserOut struct {
	ID int64 `json:"id"`
}

// UserRepo Interface for user repository.
type UserRepo interface {
	Create(context.Context, *CreateUserIn) (*CreateUserOut, error)
	Delete(context.Context, int64) (string, error)
	FindByID(context.Context, int64) (*model.User, error)
	ListAll(context.Context) ([]*model.User, error)
}
