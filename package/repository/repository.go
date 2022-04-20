package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khusainnov/rest-api/internal/user"
)

type Authorization interface {
	CreateUser(u user.User) (int, error)
}

type PostList interface {
}

type Repository struct {
	Authorization
	PostList
}

func NewService(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
