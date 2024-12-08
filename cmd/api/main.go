package main

import (
	"context"
	"fmt"
	"time"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/app"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		if err := ctx.Err(); err == nil {
			return
		}
		fmt.Printf("failed to run app: %s\n", ctx.Err())
	}()

	a, err := app.New()
	if err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}

	if err := a.Configure(ctx); err != nil {
		return fmt.Errorf("failed to configure app: %w", err)
	}

	if err := a.Run(); err != nil {
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}
