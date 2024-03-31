package main

import (
	"fmt"
	"github.com/ulugbek0217/template/pkg/config"
	"github.com/ulugbek0217/template/pkg/handlers"
	"github.com/ulugbek0217/template/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var templateList config.TemplatesConfig
	var err error = nil

	// create template cache
	templateList.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Couldn't create template cache")
	}
	// control the mode: real time or using caching system. true -> use caching, false -> real time editing
	templateList.UseCache = false

	repo := handlers.NewRepo(&templateList)
	handlers.NewHandlers(repo)

	render.SetTemplateConfig(&templateList)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting server on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
