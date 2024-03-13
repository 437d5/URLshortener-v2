package server

import (
	//"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/437d5/URLshortener-v2/handlers"
)

func newRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/urls", loadRoutes)

	return router
}

func loadRoutes(router chi.Router) {
	urlHandler := &handlers.URL{}

	router.Get("/", urlHandler.HelloHandler)
	// TODO: create url redo
	router.Post("/create", urlHandler.CreateURL)
	router.Get("/{id}", urlHandler.GetURL)
	router.Delete("/{id}", urlHandler.DeleteURL)
}
