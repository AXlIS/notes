package service

import "github.com/AXlIS/notes/internal/repository"

type Authorization interface {
	CreateUser() (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Note interface {
	Create(userID int) (int, error)
	GetAll(userID int) error
	GetByID(userID int) error
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
