package api

import (
	"demo_ecommerce/service"

	"github.com/gin-gonic/gin"
)

type UserSite struct {
	userService service.IUser
}

func NewUser(r *gin.Engine, userService service.IUser) {
	handler := &UserSite{
		userService: userService,
	}

	Group := r.Group("v1/user")
	{

	}
}
