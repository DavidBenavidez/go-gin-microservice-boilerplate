package config

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) setupRoutes() {
	s.router = chi.NewRouter()
	s.router.Use(middleware.Logger)
	s.router.Get("/", s.handleProject())
}
