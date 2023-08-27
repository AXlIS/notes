package repository

import (
	"github.com/AXlIS/notes/internal/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Note interface {
	Create(userID int, note model.Note) (int, error)
	GetAll(userID int) ([]model.Note, error)
	GetById(userID, noteID int) (model.Note, error)
}

type Repository struct {
	Authorization
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Note:          NewNotePostgres(db),
	}
}
