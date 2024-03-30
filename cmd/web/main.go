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
	templateList.TemplateCache, err = render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Couldn't create template cache")
	}

	render.SetTemplateConfig(&templateList)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting server on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
