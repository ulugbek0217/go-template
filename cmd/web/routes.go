package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ulugbek0217/template/pkg/config"
	"github.com/ulugbek0217/template/pkg/handlers"
	"net/http"
)

func routes(app *config.TemplatesConfig) http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(LoadSession)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
