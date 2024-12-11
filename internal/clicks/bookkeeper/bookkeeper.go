package bookkeeper

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/x_publisher"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain/runner"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/cfgutil"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging/lf"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xredis"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xtime"
)

type App struct {
	config Config
	logger logging.Logger
	runner *runner.Runner
}

func New() (*App, error) {
	c := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg := Config{}
	if err := cfgutil.NewLoader(*c).Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed reading config: %w", err)
	}

	app := &App{
		config: cfg,
		logger: logging.NewSLogger(), // todo: inject config
	}

	app.logger.Debug("config", lf.Any("config", cfg))

	return app, nil
}

func (a *App) Configure(ctx context.Context) error {
	redisClient, err := xredis.NewClient(ctx, a.config.Redis)
	if err != nil {
		return fmt.Errorf("failed to create redis client: %w", err)
	}

	retriever := redis_tile_storage.New(redisClient, a.config.TileStorage)

	a.runner = runner.New(
		a.config.Runner,
		x_publisher.New(),
		retriever,
		xtime.ActualProvider{},
		a.logger,
	)

	return nil
}

func (a *App) Run() error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		a.runner.Run()
	}()

	<-signalChan
	if err := a.runner.GracefulShutdown(); err != nil {
		return fmt.Errorf("failed to gracefully shutdown runner: %w", err)
	}

	return nil
}
