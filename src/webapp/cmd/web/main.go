package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
	"webapp/pkg/render"
)

const port = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// routes
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", port)
	// listen for requests
	_ = http.ListenAndServe(port, nil)
}
