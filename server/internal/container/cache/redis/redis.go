package redis_provider

import (
	"context"
	"fmt"
	"stockanalyzer/internal/container/config"

	"github.com/redis/go-redis/v9"
)

// Redis interface

type Redis struct {
	Client *redis.Client
}

// NewRedis creates a new Redis instance
func NewRedisProvider(config config.RedisConfig) *Redis {
	//opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	// opt, err := redis.ParseURL(fmt.Sprintf("%s://%s:%s@%s:%s", config.Host, config.User, config.Password, config.Host, config.Port))
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("Redis config: %+v\n", config)

	opt := &redis.Options{
		Addr:     "redis:6379",
		Password: config.Password,
	}

	client := redis.NewClient(opt)

	if client == nil {
		panic("Failed to create Redis client")

	}

	return &Redis{
		Client: client,
	}
}

// GetConnection returns the Redis client
func (r *Redis) GetConnection() interface{} {
	return r.Client
}

// Get returns the value of the key
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// Set sets the value of the key
func (r *Redis) Set(ctx context.Context, key string, value interface{}) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}
