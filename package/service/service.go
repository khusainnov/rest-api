package service

import (
	"github.com/khusainnov/rest-api/internal/user"
	"github.com/khusainnov/rest-api/package/repository"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type PostList interface {
}

type Service struct {
	Authorization
	PostList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
