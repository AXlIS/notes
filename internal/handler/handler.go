package handler

import (
	"github.com/AXlIS/notes/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", h.login)
			r.Post("/sing-up", h.singUp)
		})

		r.Route("/notes", func(r chi.Router) {
			r.Post("/", h.createNote)
			r.Get("/", h.getAllNotes)
			r.Get("/{id}", h.getNoteByID)
		})

	})

	return router
}
