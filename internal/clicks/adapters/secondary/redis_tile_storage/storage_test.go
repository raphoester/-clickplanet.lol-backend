package redis_tile_storage_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/domain"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xenvs"
	"github.com/stretchr/testify/suite"
)

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	suite.Suite
	storage interface {
		Set(ctx context.Context, tile uint32, value string) error
		Subscribe(ctx context.Context) (<-chan domain.TileUpdate, error)
		GetStateBatch(ctx context.Context, start uint32, end uint32) (map[uint32]string, error)
	}
	redis *xenvs.Redis
}

func (s *testSuite) SetupSuite() {
	var err error
	s.redis, err = xenvs.NewRedis()
	s.Require().NoError(err)

	setAndPublishOnStreamSha1 := s.redis.ScriptsMap["setAndPublishOnStream"]
	s.storage = redis_tile_storage.NewStreamStorage(s.redis.Client, setAndPublishOnStreamSha1)

	//setAndPublishSha1 := s.redis.ScriptsMap["setAndPublish"]
	//s.storage = redis_tile_storage.New(s.redis.Client, setAndPublishSha1)
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

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
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
