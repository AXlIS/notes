package service

import (
	"github.com/AXlIS/notes/internal/model"
	"github.com/AXlIS/notes/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (*Token, error)
	ParseToken(token string) (int, error)
}

type Note interface {
	Create(userID int, note model.Note) (int, error)
	GetAll(userID int) ([]model.Note, error)
	GetByID(userID, noteID int) (model.Note, error)
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Note:          NewNoteService(repos.Note),
	}
}
