package redis_tile_storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	SetAndPublishOnStreamSha1 string
}

func New(
	redis *redis.Client,
	config Config,
) *Storage {
	return &Storage{
		redis:                     redis,
		setAndPublishOnStreamSha1: config.SetAndPublishOnStreamSha1,
	}
}

const streamLabel = "tileUpdates"

type Storage struct {
	redis                     *redis.Client
	setAndPublishOnStreamSha1 string
}

func (s *Storage) Set(ctx context.Context, tile uint32, value string) error {
	_, err := s.redis.EvalSha(
		ctx,
		s.setAndPublishOnStreamSha1,
		[]string{strconv.FormatUint(uint64(tile), 10)},
		value, streamLabel,
	).Result()

	if err != nil {
		return fmt.Errorf("failed to set tile: %w", err)
	}

	return nil
}

func (s *Storage) GetStateBatch(ctx context.Context, start uint32, end uint32) (map[uint32]string, error) {
	keys := make([]string, 0, end-start+1)
	for i := start; i <= end; i++ {
		keys = append(keys, strconv.FormatUint(uint64(i), 10))
	}

	values, err := s.redis.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get tile values: %w", err)
	}

	retMap := make(map[uint32]string, len(values))
	for i, key := range keys {
		if values[i] == nil {
			continue
		}

		tile, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse tileId to int: %w", err)
		}

		retMap[uint32(tile)] = values[i].(string)
	}

	return retMap, nil
}

func (s *Storage) Subscribe(ctx context.Context) (<-chan domain.TileUpdate, error) {
	startID := "$"
	ch := make(chan domain.TileUpdate)
	errsCh := make(chan error) // TODO: do something with the errsCh
	ready := make(chan struct{})

	go func() {
		wg := sync.WaitGroup{}

		for {
			select {
			case <-ctx.Done():
				wg.Wait()
				close(ch)
				return
			default:
				xReadRes := s.redis.XRead(ctx, &redis.XReadArgs{
					Streams: []string{streamLabel, startID},
					Count:   100,
					Block:   1 * time.Second, // DO NOT SET AT 0 OTHERWISE IT WILL BLOCK FOREVER
				})

				// signal that the subscription is ready, in a select otherwise it will block the goroutine
				select {
				case ready <- struct{}{}:
				default:
				}

				streams, err := xReadRes.Result()
				if err != nil && !errors.Is(err, redis.Nil) { // wtf redis lib
					errsCh <- fmt.Errorf("failed to read stream: %w", err)
					continue
				}

				if len(streams) == 0 {
					continue
				}

				if len(streams[0].Messages) == 0 {
					continue
				}

				startID = streams[0].Messages[len(streams[0].Messages)-1].ID

				wg.Add(1)
				// handle slow consumers in a separate goroutine
				// the waitGroup will be done when the slow consumer is done, making sure the goroutine is not closed
				// because of the context finishing before the slow consumer is done
				go func() {
					defer wg.Done()
					for _, message := range streams[0].Messages {

						tileUpdate, err := xMessageToTileUpdate(message)
						if err != nil {
							errsCh <- fmt.Errorf("failed to parse message to tile update: %w", err)
							continue
						}

						select {
						case ch <- *tileUpdate:
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
	}()

	<-ready

	return ch, nil
}

// PastUpdates returns all updates that happened between now and duration ago
func (s *Storage) PastUpdates(
	ctx context.Context,
	duration time.Duration,
	now time.Time,
) ([]domain.TileUpdate, error) {
	startTime := now.Add(-duration)
	startID := fmt.Sprintf("%d-0", startTime.UnixMilli())

	response, err := s.redis.XRange(ctx, streamLabel, startID, "+").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to read Redis stream: %w", err)
	}

	updates := make([]domain.TileUpdate, 0, len(response))
	for _, message := range response {
		update, err := xMessageToTileUpdate(message)
		if err != nil {
			return nil, fmt.Errorf("failed to parse message to tile update: %w", err)
		}

		updates = append(updates, *update)
	}

	return updates, nil
}

func xMessageToTileUpdate(message redis.XMessage) (*domain.TileUpdate, error) {
	// reducing bandwidth
	// t : tile
	// n : new value
	// o : old value

	tileId, err := strconv.ParseUint(fmt.Sprint(message.Values["t"]), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tileId to int: %w", err)
	}

	return &domain.TileUpdate{
		Tile:     uint32(tileId),
		Value:    message.Values["n"].(string),
		Previous: message.Values["o"].(string),
	}, nil
}
