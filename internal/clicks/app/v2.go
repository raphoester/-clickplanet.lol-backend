package app

import (
	"context"
	"fmt"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/primary/http/clicks_controller"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/primary/http/websocket_publisher"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/in_memory_country_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/in_memory_tile_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain/click_handler_service"
)

func (a *App) configureAppV2(ctx context.Context) (*ConfigureAppResponse, error) {
	a.configurePromRegistryIfNeeded()
	a.configureHTTPFormatsIfNeeded()
	if err := a.configureRedisClientIfNeeded(ctx); err != nil {
		return nil, fmt.Errorf("failed to configure redis client: %w", err)
	}

	tilesChecker := in_memory_tile_checker.New(a.config.GameMap.MaxIndex)
	countryChecker := in_memory_country_checker.New()
	tilesStorage := redis_tile_storage.NewStreamStorage(a.redisClient, a.config.TilesStorage.SetAndPublishOnStreamSha1)

	var clickHandlerService click_handler_service.IService = click_handler_service.New(
		tilesChecker,
		tilesStorage,
		countryChecker,
	)

	//clickHandlerService, err := prom_click_handler_service.New(clickHandlerService, a.promRegistry)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create prometheus click handler service: %w", err)
	//}

	updatesCh, err := tilesStorage.Subscribe(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to tile updates: %w", err)
	}

	publisher := websocket_publisher.New(updatesCh, a.answerer)
	a.runners = append(a.runners, publisher.Run)
	controller := clicks_controller.New(
		clickHandlerService,
		tilesChecker,
		tilesStorage,
		a.answerer,
		a.reader,
	)

	return &ConfigureAppResponse{
		declareWSRoutes:  publisher.DeclareRoutes,
		declareRPCRoutes: controller.DeclareRoutes,
	}, nil

}
