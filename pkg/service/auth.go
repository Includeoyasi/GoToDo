package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Includeoyasi/todo-app"
	"github.com/Includeoyasi/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "3k5h3kfn9v8d0g8f"
	SigningKey = "1j4k5nfkd9v709v8bn9n"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return as.repo.CreateUser(user)
}

func (as *AuthService) ParseToken(Accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(Accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}
		return []byte(SigningKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type of *tokenClaims")
	}

	return claims.UserId, nil
}

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := as.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(SigningKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
