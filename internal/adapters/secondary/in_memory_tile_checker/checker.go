package in_memory_tile_checker

func New(tiles []string) *Checker {
	tilesSet := make(map[string]struct{}, len(tiles))
	for _, tile := range tiles {
		tilesSet[tile] = struct{}{}
	}
	return &Checker{
		tiles: tilesSet,
	}
}

type Checker struct {
	tiles map[string]struct{}
}

func (c *Checker) Check(tile string) bool {
	_, ok := c.tiles[tile]
	return ok
}
