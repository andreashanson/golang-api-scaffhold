package server

import (
	"net/http"

	"example.com/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Handlers *handlers.Handlers
	Server   *http.Server
}

func NewServer(h *handlers.Handlers) *Server {
	r := chi.NewRouter()

	r.Get("/status", h.GetAll())

	return &Server{Handlers: h, Server: &http.Server{
		Addr:    ":8000",
		Handler: r,
	}}
}

func (s *Server) Start() error {
	return s.Server.ListenAndServe()
}
