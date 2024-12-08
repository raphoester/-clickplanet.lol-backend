package app

import "github.com/raphoester/clickplanet.lol-backend/internal/kernel/xredis"

type Config struct {
	HTTPServer   HTTPServerConfig
	GameMap      GameMapConfig
	TilesStorage TilesStorageConfig
}

type HTTPServerConfig struct {
	BindAddress string
	Format      string
}

type GameMapConfig struct {
	MaxIndex uint32
}

type TilesStorageConfig struct {
	Redis                     xredis.Config
	SetAndPublishSha1         string
	SetAndPublishOnStreamSha1 string
}
