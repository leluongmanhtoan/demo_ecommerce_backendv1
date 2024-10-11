package model

import "github.com/uptrace/bun"

type Role struct {
	bun.BaseModel `bun:"role"`
	RoleUuid      string `json:"id" bun:"id,type:uuid,pk,notnull"`
	RoleName      string `json:"name" bun:"name,type:varchar(20),notnull"`
}
