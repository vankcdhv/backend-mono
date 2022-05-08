package mysql

import (
	"backend-mono/database/model"
	"backend-mono/database/repo"
	"context"
)

type UserDB struct {
}

func NewUserDB() (*UserDB, error) {
	return &UserDB{}, nil
}

func (u UserDB) Create(ctx context.Context, in *repo.CreateUserIn) (*repo.CreateUserOut, error) {
	return &repo.CreateUserOut{
		Message: "create user successfully",
	}, nil
}

func (u UserDB) Delete(ctx context.Context, i int64) (string, error) {
	return "Delete user successfully", nil
}

func (u UserDB) FindByID(ctx context.Context, i int64) (*model.User, error) {
	return &model.User{
		ID:        "1",
		FirstName: "Lê",
		LastName:  "Văn",
	}, nil
}

func (u UserDB) ListAll(ctx context.Context) ([]*model.User, error) {
	res := make([]*model.User, 0)
	res = append(res, &model.User{
		ID:        "1",
		FirstName: "Lê",
		LastName:  "Văn",
	})
	return res, nil
}
