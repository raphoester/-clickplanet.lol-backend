package app

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/primary/http/clicks_controller"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/primary/http/websocket_publisher"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_country_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_map_getter"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_tile_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/cfgutil"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/httpserver"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
)

func Run() error {
	logger := logging.NewSLogger()

	defer func() {
		if r := recover(); r != nil {
			logger.Error("recovered from panic", lf.Any("panic", r))
		}
	}()

	c := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg := Config{}
	if err := cfgutil.NewLoader(*c).Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed reading config: %w", err)
	}

	logger.Debug("Config loaded", lf.Any("config", cfg))

	gameMap, err := game_map.Generate(cfg.GameMap)
	if err != nil {
		return fmt.Errorf("failed to generate game map: %w", err)
	}

	answerer := httpserver.NewAnswerer(logger, httpserver.AnswerModeJSON)

	tilesChecker := in_memory_tile_checker.New(gameMap.Tiles)
	countryChecker := in_memory_country_checker.New()
	mapGetter := in_memory_map_getter.New(gameMap)
	tilesStorage := in_memory_tile_storage.New()

	router := http.NewServeMux()

	publisher := websocket_publisher.New(tilesStorage.Subscribe(), answerer)
	wsRouter := http.NewServeMux()
	wsRouter.HandleFunc("GET /listen", publisher.Subscribe)

	controller := clicks_controller.New(
		tilesChecker,
		countryChecker,
		mapGetter,
		tilesStorage,
		answerer,
		httpserver.JSONReader{},
	)

	appRouter := http.NewServeMux()
	appRouter.HandleFunc("GET /map", controller.GetMap)
	appRouter.HandleFunc("POST /click", controller.HandleClick)

	appMiddlewares := httpserver.MiddlewareStack(
		httpserver.NewLoggingMiddleware(logger),
		httpserver.CorsMiddleware,
	)

	router.Handle("/app/",
		http.StripPrefix("/app", appMiddlewares(appRouter)),
	)

	router.Handle("/ws/",
		http.StripPrefix("/ws", wsRouter),
	)

	server := http.Server{
		Addr:    cfg.HTTPServer.BindAddress,
		Handler: router,
	}

	logger.Info("Listening",
		lf.String("address", server.Addr),
	)

	go publisher.Run()

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
