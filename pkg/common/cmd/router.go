package cmd

import (
	"gihub.com/go-chi/chi/middleware"
	"github.com/go-chi/chi"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.logger)
	return r
}