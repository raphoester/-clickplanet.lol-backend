package app

type Config struct {
	GRPCServer GRPCServerConfig
}

type GRPCServerConfig struct {
	BindAddress      string
	EnableReflection bool
}
