package service

import (
	"context"
	"crypto/rand"
	"demo_ecommerce/common/model"
	"demo_ecommerce/common/response"
	"demo_ecommerce/repository"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	IUser interface {
		UserSignUp(ctx context.Context, userPost model.PostSignUp) (int, any)
		UserSignIn(ctx context.Context, userPost model.PostSignIn) (int, any)
	}

	User struct {
		jwtService JwtService
	}
)

func NewUser(jwtservice JwtService) IUser {
	return &User{
		jwtService: jwtservice,
	}
}

func (s *User) UserSignUp(ctx context.Context, userPost model.PostSignUp) (int, any) {
	userExisted, _ := repository.UserRepo.GetUserByUsername(ctx, userPost.Username)
	if userExisted != nil {
		return response.ServiceUnavailableMessage("username already exists")
	}
	salt, _ := generateSalt(16)
	hashPassword, _ := hashPassword(userPost.Password, salt)
	user := model.User{
		UserUuid:  uuid.NewString(),
		Username:  userPost.Username,
		Salt:      salt,
		Hash:      hashPassword,
		Fullname:  userPost.Firstname + " " + userPost.Lastname,
		Email:     userPost.Email,
		Address:   userPost.Address,
		RoleId:    userPost.RoleUuid,
		CreatedAt: time.Now(),
		Deleted:   false,
	}

	if err := repository.UserRepo.InsertUser(ctx, &user); err != nil {
		return response.ServiceUnavailableMessage(err.Error())
	}
	return response.Created(map[string]any{
		"user_uuid": user.UserUuid,
	})
}

func (s *User) UserSignIn(ctx context.Context, userPost model.PostSignIn) (int, any) {
	userRes, err := repository.UserRepo.GetUserByUsername(ctx, userPost.Username)

	if err != nil {
		return response.ServiceUnavailableMessage("username is not available")
	} else if userRes.Username == "" {
		return response.ServiceUnavailableMessage("username is not available")
	}
	password := userPost.Password
	hash := userRes.Hash
	salt := userRes.Salt
	err = compareHashandPassword(hash, password, salt)
	if err != nil {
		return response.ServiceUnavailableMessage("wrong password")
	}
	generatedToken := s.jwtService.GenerateToken(userRes.Username)
	return http.StatusAccepted, map[string]any{
		"status":  http.StatusText(http.StatusAccepted),
		"code":    http.StatusAccepted,
		"message": "login successful",
		"token":   generatedToken,
	}

}

// Password handler
func generateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func hashPassword(password, salt string) (string, error) {
	saltedPassword := salt + password
	hash, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(hash), nil
}

func compareHashandPassword(hash, password, salt string) error {
	saltedPassword := salt + password
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(saltedPassword))
}
