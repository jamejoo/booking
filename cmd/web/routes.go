package main

import (
	"net/http"

	"github.com/dddga/bookings/pkg/config"
	"github.com/dddga/bookings/pkg/handlers"
	"github.com/go-chi/chi/middleware"
	chi "github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {

	/*
		mux := pat.New()

		mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	*/

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
