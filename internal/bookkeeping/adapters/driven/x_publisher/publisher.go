package x_publisher

import (
	"context"

	"github.com/raphoester/clickplanet.lol-backend/internal/bookkeeping/runner"
)

func New() *Publisher {
	return &Publisher{}
}

type Publisher struct {
}

func (p *Publisher) Publish(ctx context.Context, update *runner.Update) error {
	return nil
}
