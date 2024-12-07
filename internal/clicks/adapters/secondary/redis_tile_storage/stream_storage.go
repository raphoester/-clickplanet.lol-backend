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

func NewStreamStorage(
	redis *redis.Client,
	setAndPublishOnStreamSha1 string,
) *StreamStorage {
	return &StreamStorage{
		redis:                     redis,
		setAndPublishOnStreamSha1: setAndPublishOnStreamSha1,
	}
}

const streamLabel = "tileUpdates"

type StreamStorage struct {
	redis                     *redis.Client
	setAndPublishOnStreamSha1 string
}

func (s *StreamStorage) Set(ctx context.Context, tile uint32, value string) error {
	_, err := s.redis.EvalSha(
		ctx,
		s.setAndPublishOnStreamSha1,
		[]string{strconv.FormatUint(uint64(tile), 10)},
		value, channel,
	).Result()

	if err != nil {
		return fmt.Errorf("failed to set tile: %w", err)
	}

	return nil
}

func (s *StreamStorage) GetStateBatch(ctx context.Context, start uint32, end uint32) (map[uint32]string, error) {
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

func (s *StreamStorage) Subscribe(ctx context.Context) (<-chan domain.TileUpdate, error) {
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
						// reducing bandwidth
						// t : tile
						// n : new value
						// o : old value

						tileId, err := strconv.ParseUint(fmt.Sprint(message.Values["t"]), 10, 32)
						if err != nil {
							errsCh <- fmt.Errorf("failed to parse tileId to int: %w", err)
							continue
						}

						select {
						case ch <- domain.TileUpdate{
							Tile:     uint32(tileId),
							Value:    message.Values["n"].(string),
							Previous: message.Values["o"].(string),
						}:

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
