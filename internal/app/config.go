package app

import "github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"

type Config struct {
	HTTPServer HTTPServerConfig
	GameMap    game_map.GameMapConfig
}

type HTTPServerConfig struct {
	BindAddress string
}
