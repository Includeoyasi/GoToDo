package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Includeoyasi/todo-app"
	"github.com/Includeoyasi/todo-app/pkg/repository"
)

const salt = "asgsa76as867ads"

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

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
