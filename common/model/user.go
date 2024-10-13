package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users"`
	UserUuid      string    `json:"id" bun:"id,type:uuid,pk,notnull"`
	Fullname      string    `json:"fullname" bun:"fullname,type:varchar(100),notnull"`
	Username      string    `json:"username" bun:"username,type:varchar(50),notnull"`
	Salt          string    `json:"salt" bun:"salt,type:varchar(64),notnull"`
	Hash          string    `json:"hash" bun:"hash,type:varchar(255),notnull"`
	Email         string    `json:"email" bun:"email,type:varchar(150),notnull"`
	Address       string    `json:"address" bun:"address,type:varchar(200),notnull"`
	RoleId        string    `json:"role_id" bun:"role_id,type:uuid,notnull"`
	Deleted       bool      `json:"deleted" bun:"deleted,type:boolean,notnull"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at,type:timestamp,notnull,nullzero"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at,type:timestamp,nullzero"`
}

type PutUserPassword struct {
	CurrentPassword string `json:"current_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type (
	PostSignUp struct {
		RoleUuid  string `json:"role_id"`
		Address   string `json:"address"`
		Lastname  string `json:"last_name"`
		Firstname string `json:"first_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	PostSignIn struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
