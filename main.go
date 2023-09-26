package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/YonatanGott/go-micro/app"
)

func main() {
	application := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := application.Start(ctx)
	if err != nil {
		fmt.Println("failed to start server: %w", err)
	}
}
