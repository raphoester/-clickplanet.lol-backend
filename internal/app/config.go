package app

import "github.com/raphoester/clickplanet.lol-backend/internal/pkg/redis_helper"

type Config struct {
	HTTPServer   HTTPServerConfig
	GameMap      GameMapConfig
	TilesStorage TilesStorageConfig
}

type HTTPServerConfig struct {
	BindAddress string
}

type GameMapConfig struct {
	MaxIndex uint32
}

type TilesStorageConfig struct {
	Redis             redis_helper.Config
	SetAndPublishSha1 string
}
