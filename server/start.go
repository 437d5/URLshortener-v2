package server

import (
	"net/http"
	"context"

	"github.com/437d5/URLshortener-v2/server"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	app := &App{
		router: newRoutes(),
	}

	return app
}