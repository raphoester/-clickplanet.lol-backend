package app

type Config struct {
	HTTPServer HTTPServerConfig
	GameMap    GameMapConfig
	Redis      RedisConfig
}

type HTTPServerConfig struct {
	BindAddress string
}

type GameMapConfig struct {
	MaxIndex uint32
}

type RedisConfig struct {
	Addr              string
	Password          string
	DB                int
	SetAndPublishSha1 string
}
