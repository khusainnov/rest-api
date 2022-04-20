package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/khusainnov/rest-api/internal/user"
	"github.com/khusainnov/rest-api/package/repository"
)

const salt = "5j4n34njh3b"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(u user.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)

	return s.repo.CreateUser(u)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
