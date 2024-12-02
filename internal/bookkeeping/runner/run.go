package runner

import (
	"context"
	"fmt"
)

type Update struct {
}

func (r *Runner) runOnce(ctx context.Context) error {
	update, err := r.retriever.Retrieve(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve update: %w", err)
	}

	err = r.publisher.Publish(ctx, update)
	if err != nil {
		return fmt.Errorf("failed to publish update: %w", err)
	}

	return nil
}
