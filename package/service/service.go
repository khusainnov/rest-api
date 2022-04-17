package service

import "github.com/khusainnov/rest-api/package/repository"

type Authorization interface {

}

type PostList interface {

}

type Service struct {
	Authorization
	PostList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
