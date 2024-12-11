package runner

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging/lf"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xtime"
)

type Config struct {
	Interval time.Duration
}

func New(
	cfg Config,
	publisher Publisher,
	retriever Retriever,
	timeProvider xtime.Provider,
	logger logging.Logger,
) *Runner {
	return &Runner{
		interrupt:    make(chan struct{}),
		interval:     cfg.Interval,
		logger:       logger,
		timeProvider: timeProvider,

		publisher: publisher,
		retriever: retriever,
	}
}

type Runner struct {
	interrupt    chan struct{}
	interval     time.Duration
	timeProvider xtime.Provider
	logger       logging.Logger
	publisher    Publisher
	retriever    Retriever
}

type Publisher interface {
	Publish(ctx context.Context, update *Update) error
}

type Retriever interface {
	PastUpdates(ctx context.Context, duration time.Duration, now time.Time) ([]domain.TileUpdate, error)
}

func (r *Runner) Run() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(r.interval).Do(r.runOnce)
	if err != nil {
		r.logger.Error("failed to schedule job", lf.Err(err))
		return
	}

	scheduler.StartAsync()

	<-r.interrupt

	scheduler.Stop()

	r.interrupt = nil
}

func (r *Runner) GracefulShutdown() error {
	if r.interrupt != nil {
		r.interrupt <- struct{}{}
	}
	return nil
}
