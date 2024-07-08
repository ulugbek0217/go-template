package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf adds CSRF protextion to all post requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// LoadSession Loads and saves the session in each request
func LoadSession(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
