package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kzthet194/go-course/pkg/config"
	"github.com/kzthet194/go-course/pkg/handlers"
	"github.com/kzthet194/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
