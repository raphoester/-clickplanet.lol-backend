package bookkeeper

import (
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain/runner"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xredis"
)

type Config struct {
	Runner      runner.Config
	Redis       xredis.Config
	TileStorage redis_tile_storage.Config
}
