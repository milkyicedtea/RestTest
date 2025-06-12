package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func InitRedis(ctx context.Context, cfg *Config) error {
	var err error
	redisOnce.Do(func() {
		addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
		log.Printf("Initializing Redis at %s", addr)
		redisClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: cfg.RedisPassword,
			DB:       0, // default
		})

		// test connection
		if pingErr := redisClient.Ping(ctx).Err(); pingErr != nil {
			err = fmt.Errorf("failed to connect to Redis: %w", pingErr)
			redisClient = nil
		}
	})
	return err
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		log.Panic("Redis client not initialized. Call InitRedis first.")
	}
	return redisClient
}

func CloseRedis() {
	if redisClient != nil {
		if err := redisClient.Close(); err != nil {
			log.Printf("Error closing Redis client: %v", err)
		} else {
			log.Println("Redis client closed.")
		}
	}
}
