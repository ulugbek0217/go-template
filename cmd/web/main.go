package main

import (
	"fmt"
	"github.com/ulugbek0217/template/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting server on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
