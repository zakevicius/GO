package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Handle("/static/*", FileServer())

	//mux.Get("/static/*", func (w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r)
	//})

	return mux
}

func FileServer(a ...interface{}) http.Handler {
	return http.StripPrefix("/static", http.FileServer(http.Dir("../../static/")))
}
