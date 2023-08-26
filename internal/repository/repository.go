package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
	CreateUser() (int, error)
	GetUser(username, password string) error
}

type Note interface {
	Create()
	GetAll()
	GetById()
}

type Repository struct {
	Authorization
	Note
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
