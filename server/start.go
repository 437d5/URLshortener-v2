package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	//"golang.org/x/text/cases"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func NewApp() *App {
	app := &App{
		router: newRoutes(),
		rdb: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_URL"),
			//Password: os.Getenv("REDIS_PASS"),
		}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("starting server")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

	return nil
}
