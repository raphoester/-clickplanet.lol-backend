package redis_tile_storage_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/raphoester/clickplanet.lol-backend/internal/clicks/adapters/secondary/redis_tile_storage"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xenvs"
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

	setAndPublishSha1 := s.redis.ScriptsMap["setAndPublish"]
	s.storage = redis_tile_storage.New(s.redis.Client, setAndPublishSha1)
}

func (s *testSuite) TearDownSuite() {
	s.Require().NoError(s.redis.Destroy())
}

func (s *testSuite) SetupTest() {
	s.Require().NoError(s.redis.Clean())
}

func (s *testSuite) TestSetAndPublish() {
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
			s.T().Fatal("timeout")
		case val := <-listener:
			s.T().Logf("received %v", val)
		}
	}()

	err = s.storage.Set(context.Background(), 10, "test")
	s.Require().NoError(err)

	wg.Wait()
}
