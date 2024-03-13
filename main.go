package main

import (
	"context"
	"fmt"

	"github.com/437d5/URLshortener-v2/server"
)

func main() {
	app := server.NewApp()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start server:", err)
	}
}
