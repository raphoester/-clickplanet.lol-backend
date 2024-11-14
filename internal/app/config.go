package app

import "github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"

type Config struct {
	GRPCServer GRPCServerConfig
	GameMap    game_map.GameMapConfig
}

type GRPCServerConfig struct {
	BindAddress      string
	EnableReflection bool
}
