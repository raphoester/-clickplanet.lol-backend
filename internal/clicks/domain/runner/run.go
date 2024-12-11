package runner

import (
	"context"
	"fmt"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain"
)

type Update struct {
}

func (r *Runner) runOnce(ctx context.Context) error {
	updates, err := r.retriever.PastUpdates(ctx, r.interval, r.timeProvider.Now())
	if err != nil {
		return fmt.Errorf("failed to retrieve update: %w", err)
	}

	err = r.publisher.Publish(ctx, computeUpdate(updates))
	if err != nil {
		return fmt.Errorf("failed to publish update: %w", err)
	}

	return nil
}

func computeUpdate(tileUpdates []domain.TileUpdate) *Update {
	// do something with tileUpdates
	return &Update{}
}
