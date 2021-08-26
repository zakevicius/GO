package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"webapp/pkg/config"
	"webapp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Using PAT
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// Using CHI
	mux := chi.NewRouter()

	// with middlewares
	mux.Use(middleware.Recoverer)

	// custom middleware
	// mux.Use(WriteToConsole)

	// using CSRF middleware
	mux.Use(NoSurf)

	// session loading
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
