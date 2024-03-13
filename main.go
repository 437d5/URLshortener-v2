package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/437d5/URLshortener-v2/server"
)

func main() {
	app := server.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start server:", err)
	}
}
