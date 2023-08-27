package service

import (
	"github.com/AXlIS/notes/internal/model"
	"github.com/AXlIS/notes/internal/repository"
)

type NoteService struct {
	repos repository.Note
}

func NewNoteService(repos repository.Note) *NoteService {
	return &NoteService{repos: repos}
}

func (s *NoteService) Create(userID int, note model.Note) (int, error) {
	return s.repos.Create(userID, note)
}

func (s *NoteService) GetAll(userID int) ([]model.Note, error) {
	return s.repos.GetAll(userID)
}

func (s *NoteService) GetByID(userID, noteID int) (model.Note, error) {
	return s.repos.GetById(userID, noteID)
}
