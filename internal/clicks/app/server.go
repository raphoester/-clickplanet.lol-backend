package app

import (
	"context"
	"fmt"

	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/httpserver"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/prom"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xredis"
)

func (a *App) configureHTTPFormatsIfNeeded() {
	if a.reader != nil && a.answerer != nil {
		return
	}

	format := httpserver.FormatFromString(a.config.HTTPServer.Format)
	answerer, reader := format.Build(a.logger)

	a.answerer = answerer
	a.reader = reader
}

func (a *App) configureRedisClientIfNeeded(ctx context.Context) error {
	if a.redisClient != nil {
		return nil
	}

	redisClient, err := xredis.NewClient(ctx, a.config.Redis)
	if err != nil {
		return fmt.Errorf("failed to create redis client: %w", err)
	}

	a.redisClient = redisClient
	return nil
}

func (a *App) configurePromRegistryIfNeeded() {
	if a.promRegistry != nil {
		return
	}

	a.promRegistry = prom.NewRegistry()
}
