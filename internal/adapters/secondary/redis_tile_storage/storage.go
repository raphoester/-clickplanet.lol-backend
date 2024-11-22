package redis_tile_storage

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/raphoester/clickplanet.lol-backend/internal/domain"
	"github.com/redis/go-redis/v9"
)

func New(
	redis *redis.Client,
	setAndPublishSha1 string,
) *Storage {
	return &Storage{
		redis:             redis,
		setAndPublishSha1: setAndPublishSha1,
	}
}

type Storage struct {
	redis             *redis.Client
	setAndPublishSha1 string
}

const channel = "tileUpdates"

func (s *Storage) Set(ctx context.Context, tile uint32, value string) error {
	_, err := s.redis.EvalSha(
		ctx,
		s.setAndPublishSha1,
		[]string{strconv.FormatUint(uint64(tile), 10)},
		value, channel,
	).Result()

	if err != nil {
		return fmt.Errorf("failed to set tile: %w", err)
	}

	return nil
}

func (s *Storage) GetFullState(ctx context.Context) (map[uint32]string, error) {
	iter := s.redis.Scan(ctx, 0, "*", 0).Iterator()
	retMap := make(map[uint32]string)
	for iter.Next(ctx) {
		tile, err := strconv.ParseUint(iter.Val(), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse tileId to int: %w", err)
		}

		// meh
		value, err := s.redis.Get(ctx, iter.Val()).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to get tile value: %w", err)
		}

		retMap[uint32(tile)] = value
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan tiles: %w", err)
	}

	return retMap, nil
}

func (s *Storage) Subscribe(ctx context.Context) <-chan domain.TileUpdate {
	ch := make(chan domain.TileUpdate)

	go func() {
		pubSub := s.redis.Subscribe(ctx, channel)
		defer func() { _ = pubSub.Close() }()

		for msg := range pubSub.Channel() {
			payload := make(map[string]string, 1)
			if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
				continue
			}

			for tile, value := range payload { // there should be only one key
				tileId, err := strconv.ParseUint(tile, 10, 32)
				if err != nil {
					continue
				}

				ch <- domain.TileUpdate{
					Tile:  uint32(tileId),
					Value: value,
				}
			}
		}
	}()

	return ch
}
