package in_memory_tile_checker

import "github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"

func New(tiles []game_map.Tile) *Checker {
	tilesSet := make(map[string]struct{}, len(tiles))
	for _, tile := range tiles {
		tilesSet[tile.ID()] = struct{}{}
	}
	return &Checker{
		tiles: tilesSet,
	}
}

type Checker struct {
	tiles map[string]struct{}
}

func (c *Checker) CheckTile(tile string) bool {
	_, ok := c.tiles[tile]
	return ok
}
