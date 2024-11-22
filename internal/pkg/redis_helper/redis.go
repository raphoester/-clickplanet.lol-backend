package redis_helper

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address  string
	Username string
	Password string
	DB       int
}

func NewClient(config Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
		Protocol: 2,
		PoolSize: 50,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return redisClient, nil
}
