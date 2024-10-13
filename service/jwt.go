package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *JwtService {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		panic("JWT_SECRET_KEY not set")
	}

	return &JwtService{
		secretKey: secretKey,
		issuer:    "http://codelo.life",
	}
}

func (j *JwtService) GenerateToken(UserID string) string {
	claims := &struct {
		UserID string
		jwt.StandardClaims
	}{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(), // Token hết hạn sau 1 năm
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

func (j JwtService) ValidateToken(token string) (*jwt.Token, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, fmt.Errorf("not a Bearer authorization")
	}

	keyFunc := func(t_ *jwt.Token) (any, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	}
	tokenString := strings.TrimPrefix(token, "Bearer ")
	parsedToken, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}
