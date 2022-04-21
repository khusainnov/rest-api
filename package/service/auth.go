package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/khusainnov/rest-api/internal/user"
	"github.com/khusainnov/rest-api/package/repository"
	"time"
)

const (
	salt       = "5j4n34njh3b"
	signingKey = "jh43bj2jnk234j3bh4gb32HB32k21eS"
	tokenTTL   = time.Hour * 12
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

func (s *AuthService) CreateUser(u user.User) (int, error) {
	u.Password = generatePasswordHash(u.Password)

	return s.repo.CreateUser(u)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	u, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		u.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
