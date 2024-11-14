package in_memory_tiles_storage

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
