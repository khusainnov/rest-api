package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

type PostList interface {

}

type Repository struct {
	Authorization
	PostList
}

func NewService(db *sqlx.DB) *Repository {
	return &Repository{}
}

