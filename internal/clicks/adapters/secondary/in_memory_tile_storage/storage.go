package in_memory_tile_storage

import (
	"context"
	"sync"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain"
)

type Storage struct {
	tilesMu     sync.Mutex
	tiles       map[uint32]string // might do a slice ?
	subsMu      sync.Mutex
	subscribers []chan domain.TileUpdate
}

func New() *Storage {
	return &Storage{
		tiles: make(map[uint32]string),
	}
}

func (s *Storage) Set(
	_ context.Context,
	tile uint32,
	value string,
) error {
	s.tilesMu.Lock()
	previous := s.tiles[tile]
	s.tiles[tile] = value
	s.tilesMu.Unlock()

	s.subsMu.Lock()
	defer s.subsMu.Unlock()
	for _, ch := range s.subscribers {
		go func() { // handle slow subscribers
			ch <- domain.TileUpdate{
				Tile:     tile,
				Value:    value,
				Previous: previous,
			}
		}()
	}

	return nil
}

func (s *Storage) Subscribe(_ context.Context) <-chan domain.TileUpdate {
	ch := make(chan domain.TileUpdate)
	s.subsMu.Lock()
	s.subscribers = append(s.subscribers, ch)
	s.subsMu.Unlock()
	return ch
}

func (s *Storage) GetStateBatch(
	_ context.Context,
	start uint32,
	end uint32,
) (map[uint32]string, error) {
	s.tilesMu.Lock()
	defer s.tilesMu.Unlock()
	res := make(map[uint32]string)
	for i := start; i <= end; i++ {
		if val, ok := s.tiles[i]; ok {
			res[i] = val
		}
	}
	return res, nil
}
