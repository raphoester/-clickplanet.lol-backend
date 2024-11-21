package in_memory_tile_storage

import (
	"sync"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain"
)

type Storage struct {
	mu          sync.Mutex
	tiles       map[uint32]string // might do a slice ?
	subscribers []chan domain.TileUpdate
}

func New() *Storage {
	return &Storage{
		tiles: make(map[uint32]string),
	}
}

func (s *Storage) Set(tile uint32, value string) {
	s.mu.Lock()
	s.tiles[tile] = value
	s.mu.Unlock()

	for _, ch := range s.subscribers {
		go func() { // handle slow subscribers
			ch <- domain.TileUpdate{
				Tile:  tile,
				Value: value,
			}
		}()
	}
}

func (s *Storage) Get() map[uint32]string {
	return s.tiles // ownership leak but copying the map would be too expensive
}

func (s *Storage) Subscribe() <-chan domain.TileUpdate {
	ch := make(chan domain.TileUpdate)
	s.subscribers = append(s.subscribers, ch)
	return ch
}
