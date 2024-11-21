package app

type Config struct {
	HTTPServer HTTPServerConfig
	GameMap    GameMapConfig
}

type HTTPServerConfig struct {
	BindAddress string
}

type GameMapConfig struct {
	MaxIndex uint32
}
