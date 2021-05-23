package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/asiaCoder/todo-app/pkg/model"
	"github.com/asiaCoder/todo-app/pkg/repository"
)

const salt = "weef82g87f23b2_89238f2"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatedPasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatedPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
