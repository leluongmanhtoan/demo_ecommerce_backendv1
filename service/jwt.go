package service

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *jwtService {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		panic("JWT_SECRET_KEY not set")
	}

	return &jwtService{
		secretKey: secretKey,
		issuer:    "http://codelo.life",
	}
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &struct {
		UserID string
		jwt.StandardClaims
	}{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(), // Token hết hạn sau 1 năm
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}
