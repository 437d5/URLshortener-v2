package server

import (
	//"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/437d5/URLshortener-v2/db"
	"github.com/437d5/URLshortener-v2/handlers"
)

func (a *App) newRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/urls", a.loadRoutes)

	a.router = router
}

func (a *App) loadRoutes(router chi.Router) {
	urlHandler := &handlers.URL{
		Repo: &db.RedisRepo{
			Client: a.rdb,
		},
	}

	router.Get("/", urlHandler.HelloHandler)
	router.Post("/", urlHandler.CreateURL)
	router.Get("/{token}", urlHandler.GetURL)
}
