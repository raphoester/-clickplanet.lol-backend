package app

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/primary/http/clicks_controller"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/primary/http/websocket_publisher"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_country_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_tile_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/cfgutil"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/httpserver"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
	"github.com/redis/go-redis/v9"
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
	answerer := httpserver.NewAnswerer(a.logger, httpserver.AnswerModeBinary)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     a.config.Redis.Addr,
		Password: a.config.Redis.Password,
		DB:       a.config.Redis.DB,
		Protocol: 2,
		PoolSize: 50,
	})

	tilesChecker := in_memory_tile_checker.New(a.config.GameMap.MaxIndex)
	countryChecker := in_memory_country_checker.New()
	//tilesStorage := in_memory_tile_storage.New()
	tilesStorage := redis_tile_storage.New(redisClient, a.config.Redis.SetAndPublishSha1)

	updatesCh := tilesStorage.Subscribe(context.Background())
	a.publisher = websocket_publisher.New(updatesCh, answerer)
	a.controller = clicks_controller.New(
		tilesChecker,
		countryChecker,
		tilesStorage,
		answerer,
		httpserver.ProtoReader{},
	)

	a.declareRoutes()

	return nil
}

func (a *App) declareRoutes() {
	router := http.NewServeMux()
	wsRouter := http.NewServeMux()
	wsRouter.HandleFunc("GET /listen", a.publisher.Subscribe)

	appRouter := http.NewServeMux()
	appRouter.HandleFunc("GET /map-density", a.controller.GetMapDensity)
	appRouter.HandleFunc("POST /click", a.controller.HandleClick)
	appRouter.HandleFunc("GET /ownerships", a.controller.GetOwnerships)

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
