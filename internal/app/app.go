package app

import (
	"flag"
	"fmt"
	"net"

	clicksv1 "github.com/raphoester/clickplanet.lol-backend/generated/proto/clicks/v1"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/primary/grpc/clicks_controller"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_country_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/adapters/secondary/in_memory_tile_checker"
	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/cfgutil"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/grpc_stack"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_stack.LoggingUnaryInterceptor(logger)),
		grpc.StreamInterceptor(grpc_stack.LoggingStreamInterceptor(logger)),
	)

	if cfg.GRPCServer.EnableReflection {
		reflection.Register(server)
	}

	gameMap := game_map.New(cfg.GameMap)

	tilesChecker := in_memory_tile_checker.New(gameMap.Tiles)
	countryChecker := in_memory_country_checker.New()
	
	controller := clicks_controller.New(
		tilesChecker,
		countryChecker,
	)

	clicksv1.RegisterClicksServer(server, controller)

	listener, err := net.Listen("tcp", cfg.GRPCServer.BindAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on given bind address: %w", err)
	}

	logger.Info("Listening",
		lf.String("address", listener.Addr().String()),
	)

	err = server.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed serving: %w", err)
	}

	return nil
}
