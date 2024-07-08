package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Loaded the page")
		next.ServeHTTP(w, r)
	})
}

func CSRFTokenMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
