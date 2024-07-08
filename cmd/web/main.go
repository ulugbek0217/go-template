package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ulugbek0217/template/pkg/config"
	"github.com/ulugbek0217/template/pkg/handlers"
	"github.com/ulugbek0217/template/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.TemplatesConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false
	var err error

	session = scs.New()
	session.Lifetime = 2 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	// create template cache
	app.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Couldn't create template cache")
	}
	// control the mode: real time or using caching system. true -> use caching, false -> real time editing
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.SetTemplateConfig(&app)

	fmt.Printf("Starting server on %s", portNumber)
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = server.ListenAndServe()
	log.Fatal(err)

}
