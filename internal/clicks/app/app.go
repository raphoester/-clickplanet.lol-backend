package app

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/prom"
	"github.com/redis/go-redis/v9"

	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/cfgutil"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/httpserver"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/logging/lf"
)

type App struct {
	config Config
	logger logging.Logger
	server *http.Server

	runners []func()

	answerer *httpserver.Answerer
	reader   httpserver.Reader

	redisClient  *redis.Client
	promRegistry *prometheus.Registry
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
	appV2, err := a.configureAppV2(ctx)
	if err != nil {
		return fmt.Errorf("failed to configure app v2: %w", err)
	}

	rpcMiddlewares := httpserver.MiddlewareStack(
		httpserver.NewLoggingMiddleware(a.logger),
		httpserver.IPReaderMiddleware,
		httpserver.CorsMiddleware,
	)

	wsMiddlewares := httpserver.MiddlewareStack(
		//httpserver.NewLoggingMiddleware(a.logger), // TODO: fix hijacker problem
		httpserver.IPReaderMiddleware,
	)

	router := http.NewServeMux()

	appV2RPCRouter := http.NewServeMux()
	appV2.declareRPCRoutes(appV2RPCRouter)
	router.Handle("/v2/rpc/", http.StripPrefix("/v2/rpc", rpcMiddlewares(appV2RPCRouter)))

	appV2WSRouter := http.NewServeMux()
	appV2.declareWSRoutes(appV2WSRouter)
	router.Handle("/v2/ws/", http.StripPrefix("/v2/ws", wsMiddlewares(appV2WSRouter)))

	a.declarePrometheusRoutes(router)

	a.server = &http.Server{
		Addr:    a.config.HTTPServer.BindAddress,
		Handler: router,
	}

	return nil
}

func (a *App) declarePrometheusRoutes(router *http.ServeMux) {
	promRouter := http.NewServeMux()

	a.configurePromRegistryIfNeeded()
	middlewareStack := httpserver.MiddlewareStack(
		httpserver.NewLoggingMiddleware(a.logger),
	)

	promHandler := prom.HandlerForRegistry(a.promRegistry)
	promRouter.HandleFunc("GET /", promHandler.ServeHTTP)

	router.Handle("/metrics", middlewareStack(promRouter))
}

type ConfigureAppResponse struct {
	declareWSRoutes  func(mux *http.ServeMux)
	declareRPCRoutes func(mux *http.ServeMux)
}
