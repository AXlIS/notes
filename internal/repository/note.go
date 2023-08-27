package repository

import (
	"fmt"
	"github.com/AXlIS/notes/internal/model"
	"github.com/jmoiron/sqlx"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{
		db: db,
	}
}

func (r *NotePostgres) Create(userID int, note model.Note) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, text, user_id) VALUES ($1, $2, $3) RETURNING id", notesTable)

	row := r.db.QueryRow(query, note.Title, note.Text, userID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *NotePostgres) GetAll(userID int) ([]model.Note, error) {
	var noteList []model.Note

	query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE user_id = $1", notesTable)
	err := r.db.Select(&noteList, query, userID)

	return noteList, err
}

func (r *NotePostgres) GetById(userID, noteID int) (model.Note, error) {
	var note model.Note

	query := fmt.Sprintf("SELECT id, title, text FROM %s WHERE user_id = $1 and id = $2", notesTable)
	err := r.db.Get(&note, query, userID, noteID)

	return note, err
}
