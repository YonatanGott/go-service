package main

import (
	"context"
	"fmt"

	"github.com/YonatanGott/go-micro/app"
)

func main() {
	application := app.New()
	err := application.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start server: %w", err)
	}
}
