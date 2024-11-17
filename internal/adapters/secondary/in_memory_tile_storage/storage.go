package in_memory_tile_storage

import (
	"github.com/raphoester/clickplanet.lol-backend/internal/domain/game_map"
)

type Storage struct {
	tiles       map[string]string
	subscribers []chan game_map.TileUpdate
}

func New() *Storage {
	return &Storage{
		tiles: make(map[string]string),
	}
}

func (s *Storage) Set(tile string, value string) {
	s.tiles[tile] = value
	for _, ch := range s.subscribers {
		go func() { // handle slow subscribers
			ch <- game_map.TileUpdate{
				Tile:  tile,
				Value: value,
			}
		}()
	}
}

func (s *Storage) Get() map[string]string {
	return s.tiles // ownership leak but copying the map would be too expensive
}

func (s *Storage) Subscribe() <-chan game_map.TileUpdate {
	ch := make(chan game_map.TileUpdate)
	s.subscribers = append(s.subscribers, ch)
	return ch
}
