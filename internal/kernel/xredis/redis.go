package xredis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address  string
	Username string
	Password string
	DB       int
	TLS      bool
}

func NewClient(ctx context.Context, config Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
		Protocol: 2,
		PoolSize: 50,
		TLSConfig: func() *tls.Config { // weird but needs nil pointer for no tls
			if config.TLS {
				return &tls.Config{
					InsecureSkipVerify: false,
				}
			}
			return nil
		}(),
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return redisClient, nil
}
