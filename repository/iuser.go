package repository

import (
	"context"
	"demo_ecommerce/common/model"
)

type IUser interface {
	InsertUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

var UserRepo IUser
