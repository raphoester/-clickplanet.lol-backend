package xenvs

import (
	"context"
	"fmt"

	"github.com/ory/dockertest"
	"github.com/raphoester/clickplanet.lol-backend/internal/kernel/xredis"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/afero"
)

type Redis struct {
	Client     *redis.Client
	ScriptsMap map[string]string
	pool       *dockertest.Pool
	container  *dockertest.Resource
}

func (p *Redis) Destroy() error {
	if err := p.Client.Close(); err != nil {
		return err
	}
	return p.pool.Purge(p.container)
}

func (p *Redis) Clean() error {
	return p.Client.FlushAll(context.Background()).Err()
}

func NewRedis() (*Redis, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("failed creating dockertest pool: %w", err)
	}

	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "redis",
		Tag:        "7.4.1-alpine3.20",
		Env:        []string{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to run redis container: %w", err)
	}

	port := container.GetPort("6379/tcp")

	var redisClient *redis.Client
	err = pool.Retry(func() error {
		redisClient, err = xredis.NewClient(context.Background(), xredis.Config{
			Address:  fmt.Sprintf("localhost:%s", port),
			Username: "",
			Password: "",
			DB:       0,
			TLS:      false,
		})

		if err != nil {
			return fmt.Errorf("failed to connect to mongo: %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	scriptsMap, err := xredis.NewScriptLoader(
		redisClient,
		afero.NewOsFs(),
		"/static",
	).Load(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to load scripts: %w", err)
	}

	return &Redis{
		ScriptsMap: scriptsMap,
		Client:     redisClient,
		pool:       pool,
		container:  container,
	}, nil
}
