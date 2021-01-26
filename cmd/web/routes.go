package main

import (
	"github.com/go-chi/chi/middleware"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mariojackson/go-course/pkg/config"
	"github.com/mariojackson/go-course/pkg/handlers"
)

func routes(app *config.App) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
