package mysql

import (
	"backend-mono/core/logger"
	"backend-mono/database/model"
	"backend-mono/database/repo"
	"context"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	table      string
	connection *sqlx.DB
}

func NewUserDB() (*UserDB, error) {
	db, err := sqlx.Open("mysql", "dev:dev@tcp(127.0.0.1:3306)/lms")
	if err != nil {
		return nil, err
	}
	return &UserDB{
		table:      "users",
		connection: db,
	}, nil
}

func (u *UserDB) Close() {
	err := u.connection.Close()
	if err != nil {
		return
	}
}

func (u *UserDB) Create(ctx context.Context, in *repo.CreateUserIn) (*repo.CreateUserOut, error) {
	return &repo.CreateUserOut{
		Message: "create user successfully",
	}, nil
}

func (u *UserDB) Delete(ctx context.Context, i int64) (string, error) {
	return "Delete user successfully", nil
}

func (u *UserDB) FindByID(ctx context.Context, i int64) (*model.User, error) {
	ctxLogger := logger.NewContextLog(ctx)
	var result []*model.User

	db := sq.Select("*").From(u.table).Where(sq.Eq{"id": i})
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query - error: %s", err.Error())
		return nil, err
	}
	err = u.connection.Select(&result, query, arg...)
	if err != nil {
		ctxLogger.Errorf("Failed while build query - error: %s", err.Error())
		return nil, err
	}
	if err != nil {
		ctxLogger.Errorf("Failed while parse result", err.Error())
		return nil, err
	}
	return result[0], nil
}

func (u *UserDB) ListAll(ctx context.Context) ([]*model.User, error) {
	res := make([]*model.User, 0)
	res = append(res, &model.User{
		ID:        "1",
		FirstName: "Lê",
		LastName:  "Văn",
	})
	return res, nil
}
