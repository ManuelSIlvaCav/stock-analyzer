package container

import (
	"stockanalyzer/internal/container/config"
	"stockanalyzer/internal/container/postgres"
)

type Container struct {
	PostgresSQL postgres.PostgresSQL
	Config      *config.Config
}

func NewContainer() *Container {
	config := config.NewConfig()

	return &Container{
		PostgresSQL: postgres.NewPostgresSQL(config.Postgres),
		Config:      config,
	}
}

func (c *Container) GetConfig() *config.Config {
	return c.Config
}
