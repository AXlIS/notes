package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AXlIS/notes/internal/model"
	"github.com/lib/pq"
	"io"
	"net/http"
)

func (h *Handler) singUp(w http.ResponseWriter, r *http.Request) {
	var input model.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	if err = json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err = h.services.Authorization.CreateUser(input)
	if err != nil {
		pgErr, ok := err.(*pq.Error)
		if ok {
			if pgErr.Code == "23505" {
				newErrorResponse(w, http.StatusInternalServerError, "user with this username already exists")
				return
			}
		}
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var input loginInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	if err = json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			newErrorResponse(w, http.StatusUnauthorized, "invalid username or password")
			return
		}

		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("error building the response, %v", err))
		return
	}
}
