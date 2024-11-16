package in_memory_tile_storage

type Storage struct {
	tiles map[string]string
}

func New() *Storage {
	return &Storage{
		tiles: make(map[string]string),
	}
}

func (s *Storage) Set(tile string, value string) {
	s.tiles[tile] = value
}

func (s *Storage) Get() map[string]string {
	return s.tiles // ownership leak but copying the map would be too expensive
}
