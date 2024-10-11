package auth

import (
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	AuthMiddleware() gin.HandlerFunc
}

var AuthMdw IAuthMiddleware

func AuthMiddleware() gin.HandlerFunc {
	return AuthMdw.AuthMiddleware()
}
