package runner

import (
	"context"
	"time"

	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging/lf"
)

type Config struct {
	Interval time.Duration
}

func New(
	cfg Config,
	publisher Publisher,
	retriever Retriever,
	logger logging.Logger,
) *Runner {
	return &Runner{
		interrupt: make(chan struct{}),
		interval:  cfg.Interval,
		logger:    logger,

		publisher: publisher,
		retriever: retriever,
	}
}

type Runner struct {
	interrupt chan struct{}
	interval  time.Duration
	logger    logging.Logger
	publisher Publisher
	retriever Retriever
}

type Publisher interface {
	Publish(ctx context.Context, update *Update) error
}

type Retriever interface {
	Retrieve(ctx context.Context) (*Update, error)
}

func (r *Runner) Run() {
	defer func() {
		if p := recover(); p != nil {
			r.logger.Error("recovered from panic", lf.Any("panic", p))
		}
	}()
	ctx := context.Background()
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()
	for {
		select {
		case <-r.interrupt:
			r.interrupt = nil
			return
		case <-ticker.C:
			r.logger.Debug("running once")
			if err := r.runOnce(ctx); err != nil {
				r.logger.Error("failed to run once", lf.Err(err))
			}
		}
	}
}

func (r *Runner) GracefulShutdown() error {
	if r.interrupt != nil {
		r.interrupt <- struct{}{}
	}
	return nil
}
