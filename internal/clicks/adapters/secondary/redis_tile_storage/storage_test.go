package redis_tile_storage_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xenvs"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite
	storage *redis_tile_storage.Storage
	redis   *xenvs.Redis
}

func (s *testSuite) SetupSuite() {
	var err error
	s.redis, err = xenvs.NewRedis()
	s.Require().NoError(err)

	setAndPublishOnStreamSha1 := s.redis.ScriptsMap["setAndPublishOnStream"]
	s.storage = redis_tile_storage.New(s.redis.Client, redis_tile_storage.Config{SetAndPublishOnStreamSha1: setAndPublishOnStreamSha1})
}

func (s *testSuite) TearDownSuite() {
	s.Require().NoError(s.redis.Destroy())
}

func (s *testSuite) SetupTest() {
	s.Require().NoError(s.redis.Clean())
}

func (s *testSuite) TestSetAndPublish() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	listener, err := s.storage.Subscribe(ctx)
	s.Require().NoError(err)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			s.T().Errorf("timeout")
		case val := <-listener:
			s.Assert().Equal("fr", val.Value)
			s.Assert().Equal(uint32(10), val.Tile)
			s.Assert().Equal("", val.Previous)
		}
	}()

	err = s.storage.Set(context.Background(), 10, "fr")
	s.Require().NoError(err)

	wg.Wait()
}

func (s *testSuite) TestSetAndPublishWithOverride() {
	previousValue := "us"
	newValue := "fr"

	err := s.storage.Set(context.Background(), 10, previousValue)
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	listener, err := s.storage.Subscribe(ctx)
	s.Require().NoError(err)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			s.T().Errorf("timeout")
		case val := <-listener:
			s.Assert().Equal(newValue, val.Value)
			s.Assert().Equal(uint32(10), val.Tile)
			s.Assert().Equal(previousValue, val.Previous)
		}
	}()

	err = s.storage.Set(context.Background(), 10, newValue)
	s.Require().NoError(err)

	wg.Wait()
}

func (s *testSuite) TestSetAndPublishWithOverrideAndNoChange() {
	constantValue := "fr"

	err := s.storage.Set(context.Background(), 10, constantValue)
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	listener, err := s.storage.Subscribe(ctx)
	s.Require().NoError(err)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			s.T().Logf("as expected, no message was received")
		case val := <-listener:
			s.T().Errorf("unexpected value %v", val)
		}
	}()

	err = s.storage.Set(context.Background(), 10, constantValue)
	s.Require().NoError(err)

	wg.Wait()
}

func (s *testSuite) TestSetAndPublishWithALotOfConcurrentMessages() {
	constantValue := "fr"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	listener, err := s.storage.Subscribe(ctx)
	s.Require().NoError(err)

	rcvMap := make(map[uint32]struct{})
	rcvMapMu := sync.RWMutex{}

	wgRcv := sync.WaitGroup{}
	for i := 1; i <= 100_000; i++ {
		wgRcv.Add(1)
		go func() {
			defer wgRcv.Done()
			select {
			case <-ctx.Done():
				s.T().Errorf("timeout")
			case value := <-listener:
				rcvMapMu.RLock()
				s.Assert().Equal(constantValue, value.Value)
				_, ok := rcvMap[value.Tile]
				rcvMapMu.RUnlock()

				s.Assert().False(ok)

				rcvMapMu.Lock()
				rcvMap[value.Tile] = struct{}{}
				rcvMapMu.Unlock()
			}
		}()
	}

	wgSend := sync.WaitGroup{}
	for i := uint32(1); i <= 100_000; i++ {
		wgSend.Add(1)
		go func() {
			defer wgSend.Done()
			err := s.storage.Set(context.Background(), i, constantValue)
			s.Require().NoError(err)
		}()
	}

	wgSend.Wait()
	wgRcv.Wait()
	s.Assert().Equal(100_000, len(rcvMap))
}

func (s *testSuite) TestGetStateByBatch() {
	constantValue := "fr"

	err := s.storage.Set(context.Background(), 10, constantValue)
	s.Require().NoError(err)

	err = s.storage.Set(context.Background(), 20, constantValue)
	s.Require().NoError(err)

	err = s.storage.Set(context.Background(), 30, constantValue)
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	state, err := s.storage.GetStateBatch(ctx, 10, 30) // bounds are inclusive
	s.Require().NoError(err)

	s.Assert().Equal(3, len(state))
	s.Assert().Equal(constantValue, state[10])
	s.Assert().Equal(constantValue, state[20])
	s.Assert().Equal(constantValue, state[30])
}

func (s *testSuite) TestPastUpdates() {
	xAdd := func(t time.Time, i int) error {
		return s.redis.Client.XAdd(context.Background(), &redis.XAddArgs{
			Stream: "tileUpdates", // stream name corresponds in ./stream_storage.go
			ID:     fmt.Sprintf("%d-%d", t.UnixMilli(), i),
			Values: []string{"t", fmt.Sprintf("%d", i), "n", "fr", "o", "us"},
		}).Err()
	}

	xAdd1Time := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // 00:00:00

	err := xAdd(xAdd1Time, 10)
	s.Require().NoError(err)

	xAdd2Time := xAdd1Time.Add(45 * time.Minute) // 00:45:00
	err = xAdd(xAdd2Time, 11)

	xAdd3Time := xAdd2Time.Add(10 * time.Minute) // 00:55:00
	err = xAdd(xAdd3Time, 12)

	queryTime := xAdd3Time.Add(30 * time.Minute) // 01:25:00

	pastUpdates, err := s.storage.PastUpdates(context.Background(), 1*time.Hour, queryTime)
	s.Require().NoError(err)
	s.Assert().Equal(2, len(pastUpdates))

	s.Assert().Equal(uint32(11), pastUpdates[0].Tile)
	s.Assert().Equal("fr", pastUpdates[0].Value)

	s.Assert().Equal(uint32(12), pastUpdates[1].Tile)
	s.Assert().Equal("fr", pastUpdates[1].Value)
}
