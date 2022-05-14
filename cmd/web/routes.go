package main

import (
	"net/http"

	"github.com/firstneverrest/go-web-app/internal/config"
	"github.com/firstneverrest/go-web-app/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter() // multiplexer

	mux.Use(middleware.Recoverer) // recover from panics without crashing the server
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Post("/coffee", handlers.Repo.BuyCoffee)

	// enable static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
