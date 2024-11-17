package main

import (
	"fmt"

	"github.com/raphoester/clickplanet.lol-backend/internal/app"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	a, err := app.New()
	if err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}

	if err := a.Configure(); err != nil {
		return fmt.Errorf("failed to configure app: %w", err)
	}

	if err := a.Run(); err != nil {
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}
