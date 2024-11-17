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

type App struct {
	config Config
	logger logging.Logger
	server *http.Server

	publisher  *websocket_publisher.Publisher
	controller *clicks_controller.Controller
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

	return app, nil
}

func (a *App) Configure() error {
	gameMap, err := game_map.Generate(a.config.GameMap)
	if err != nil {
		return fmt.Errorf("failed to generate game map: %w", err)
	}

	answerer := httpserver.NewAnswerer(a.logger, httpserver.AnswerModeJSON)

	tilesChecker := in_memory_tile_checker.New(gameMap.Tiles)
	countryChecker := in_memory_country_checker.New()
	mapGetter := in_memory_map_getter.New(gameMap)
	tilesStorage := in_memory_tile_storage.New()

	a.publisher = websocket_publisher.New(tilesStorage.Subscribe(), answerer)

	a.controller = clicks_controller.New(
		tilesChecker,
		countryChecker,
		mapGetter,
		tilesStorage,
		answerer,
		httpserver.JSONReader{},
	)

	a.declareRoutes()

	return nil
}

func (a *App) declareRoutes() {
	router := http.NewServeMux()
	wsRouter := http.NewServeMux()
	wsRouter.HandleFunc("GET /listen", a.publisher.Subscribe)

	appRouter := http.NewServeMux()
	appRouter.HandleFunc("GET /map", a.controller.GetMap)
	appRouter.HandleFunc("POST /click", a.controller.HandleClick)
	appRouter.HandleFunc("GET /bindings", a.controller.GetBindings)

	appMiddlewares := httpserver.MiddlewareStack(
		httpserver.NewLoggingMiddleware(a.logger),
		httpserver.CorsMiddleware,
	)

	router.Handle("/app/",
		http.StripPrefix("/app", appMiddlewares(appRouter)),
	)

	router.Handle("/ws/",
		http.StripPrefix("/ws", wsRouter), // don't add middleware to websockets, causes errors
	)

	a.server = &http.Server{
		Addr:    a.config.HTTPServer.BindAddress,
		Handler: router,
	}
}

func (a *App) Run() error {
	defer func() {
		if r := recover(); r != nil {
			a.logger.Error("recovered from panic", lf.Any("panic", r))
		}
	}()

	a.logger.Info("Listening",
		lf.String("address", a.server.Addr),
	)

	go a.publisher.Run()
	if err := a.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
