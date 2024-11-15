package in_memory_map_getter

import "github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"

func New(
	theMap *game_map.GameMap,
) *Getter {
	return &Getter{
		theMap: theMap,
	}
}

type Getter struct {
	theMap *game_map.GameMap
}

func (c *Getter) GetMap() *game_map.GameMap {
	return c.theMap // ownership leak but the object is super heavy
}
