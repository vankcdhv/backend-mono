package mysql

import (
	"backend-mono/cmd/config"
	"backend-mono/cmd/database/model"
	"backend-mono/cmd/database/repo"
	"backend-mono/core/logger"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	table      string
	connection *sqlx.DB
}

func NewUserDB(c config.Config) (*UserDB, error) {
	db, err := sqlx.Open(c.Database.Driver, c.Database.Source)
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
	ctxLogger := logger.NewContextLog(ctx)
	db := sq.Insert(u.table).Columns(GetListColumn(in)...).Values(GetListValues(in)...)
	query, arg, err := db.ToSql()
	if err != nil {
		ctxLogger.Errorf("Failed while build query - error: %s", err.Error())
		return nil, err
	}
	result := u.connection.MustExec(query, arg...)
	insertedID, err := result.LastInsertId()
	if err != nil {
		ctxLogger.Errorf("Failed while insert user to database: %s", err.Error())
		return nil, err
	}
	return &repo.CreateUserOut{
		ID: insertedID,
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
		if err == sql.ErrNoRows {
			ctxLogger.Errorf("Don't have user with ID %d - error: %s", i, err.Error())
		}
		ctxLogger.Errorf("Failed while select user by ID %d - error: %s", i, err.Error())
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
