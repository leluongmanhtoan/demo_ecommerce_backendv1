package api

import (
	"demo_ecommerce/common/model"
	"demo_ecommerce/service"

	"github.com/gin-gonic/gin"
	//authMdw "demo_ecommerce/middleware/auth"
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
		Group.POST("signup", handler.PostUserSignUp)
		Group.POST("signin", handler.PostUserSignIn)

	}
}

func (h *UserSite) PostUserSignUp(c *gin.Context) {
	user := model.PostSignUp{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(200, "ok")
	}
	code, result := h.userService.UserSignUp(c, user)
	c.JSON(code, result)
}

func (h *UserSite) PostUserSignIn(c *gin.Context) {
	authinfo := model.PostSignIn{}
	if err := c.BindJSON(&authinfo); err != nil {
		c.JSON(200, "ok")
	}
	h.userService.UserSignIn(c, authinfo)
}
