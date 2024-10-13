package api

import (
	"demo_ecommerce/common/model"
	"demo_ecommerce/common/response"
	"demo_ecommerce/service"
	"net/http"

	authMdw "demo_ecommerce/middleware/auth"

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
		Group.POST("signup", handler.PostUserSignUp)
		Group.POST("signin", handler.PostUserSignIn)
		Group.GET("test", authMdw.AuthMdw.AuthMiddleware(), handler.TestAPIAuth)
	}
}

func (h *UserSite) PostUserSignUp(c *gin.Context) {
	user := model.PostSignUp{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(response.ServiceUnavailableMessage(err.Error()))
		return
	}
	code, result := h.userService.UserSignUp(c, user)
	c.JSON(code, result)
}

func (h *UserSite) PostUserSignIn(c *gin.Context) {
	authinfo := model.PostSignIn{}
	if err := c.BindJSON(&authinfo); err != nil {
		c.JSON(response.ServiceUnavailableMessage(err.Error()))
		return
	}
	code, result := h.userService.UserSignIn(c, authinfo)
	c.JSON(code, result)
}

func (h *UserSite) TestAPIAuth(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"code":    http.StatusOK,
		"message": "Authorize Successfully",
	})
}
