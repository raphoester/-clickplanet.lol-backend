package app

import (
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xredis"
)

type Config struct {
	HTTPServer   HTTPServerConfig
	GameMap      GameMapConfig
	Redis        xredis.Config
	TilesStorage redis_tile_storage.Config
}

type HTTPServerConfig struct {
	BindAddress string
	Format      string
}

type GameMapConfig struct {
	MaxIndex uint32
}
