package container

import (
	"stockanalyzer/internal/container/cache"
	redis_provider "stockanalyzer/internal/container/cache/redis"
	"stockanalyzer/internal/container/config"
	"stockanalyzer/internal/container/postgres"
)

type Container struct {
	PostgresSQL postgres.PostgresSQL
	Config      *config.Config
	Cache       cache.Cache
}

func NewContainer() *Container {
	config := config.NewConfig()

	return &Container{
		PostgresSQL: postgres.NewPostgresSQL(config.Postgres),
		Cache:       redis_provider.NewRedisProvider(config.Redis),
		Config:      config,
	}
}

func (c *Container) GetConfig() *config.Config {
	return c.Config
}
