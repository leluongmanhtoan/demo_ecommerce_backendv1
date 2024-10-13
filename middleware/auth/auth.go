package auth

import (
	"demo_ecommerce/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	AuthMiddleware() gin.HandlerFunc
}

type AuthenticationMiddleware struct {
	jwtService service.JwtService
}

var AuthMdw IAuthMiddleware

func NewAuthMiddleware(jwtservice service.JwtService) IAuthMiddleware {
	return &AuthenticationMiddleware{
		jwtService: jwtservice,
	}
}

func (auth *AuthenticationMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(
				http.StatusUnauthorized,
				map[string]any{
					"status": http.StatusText(http.StatusUnauthorized),
					"code":   http.StatusUnauthorized,
				},
			)
			c.Abort()
			return
		}
		token, err := auth.jwtService.ValidateToken(authHeader)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				map[string]any{
					"status":  http.StatusText(http.StatusUnauthorized),
					"code":    http.StatusUnauthorized,
					"message": "token is not valid",
				},
			)
			c.Abort()
		}
		claims := token.Claims.(jwt.MapClaims)
		log.Println("Claim[user_id]: ", claims["user_id"])
		log.Println("Claim[issuer] :", claims["issuer"])

	}
}
