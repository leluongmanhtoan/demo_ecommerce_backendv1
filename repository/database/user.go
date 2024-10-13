package database

import (
	"context"
	"database/sql"
	"demo_ecommerce/common/model"
	"demo_ecommerce/repository"
	"errors"
)

type User struct{}

func NewUser() repository.IUser {
	repo := &User{}
	return repo
}

func (r *User) InsertUser(ctx context.Context, user *model.User) error {
	query := repository.SqlClient.GetDB().NewInsert().Model(user)
	res, err := query.Exec(ctx)
	if err != nil {
		return err
	} else if affected, _ := res.RowsAffected(); affected != 1 {
		return errors.New("register failed")
	}
	return nil
}

func (r *User) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	query := repository.SqlClient.GetDB().NewSelect().
		Model(user).
		ColumnExpr("*").
		Where("username = ?", username).
		Limit(1)
	err := query.Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
