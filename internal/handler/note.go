package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AXlIS/notes/internal/model"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.Note
	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	if err = json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.Corrector.ValidateText(&input); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.services.Note.Create(userID, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getNoteByID(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	noteID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid note id param")
		return
	}

	note, err := h.services.Note.GetByID(userID, noteID)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			newErrorResponse(w, http.StatusNotFound, "note not found")
			return
		}

		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(note); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error building the response, %v", err))
		return
	}
}

func (h *Handler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	notes, err := h.services.Note.GetAll(userID)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error building the response, %v", err))
		return
	}
}
