package app

import (
	"fmt"

	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging/lf"
)

func (a *App) Run() error {
	a.logger.Info("Listening",
		lf.String("address", a.server.Addr),
	)

	for _, runner := range a.runners {
		go runner()
	}

	if err := a.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
