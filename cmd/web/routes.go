package main

import (
	"github.com/bmizerany/pat"
	"github.com/ulugbek0217/template/pkg/config"
	"github.com/ulugbek0217/template/pkg/handlers"
	"net/http"
)

func routes(app *config.TemplatesConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
