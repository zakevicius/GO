package main

import (
	"fmt"
	"net/http"
	"webapp/pkg/handlers"
)

const port = ":8080"

// main is the main application function
func main() {
	// routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", port)
	// listen for requests
	_ = http.ListenAndServe(port, nil)
}

