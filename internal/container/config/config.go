package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type ExternalAPIConfig struct {
	FmpApiKey string
}

type Config struct {
	Env               string
	Port              string
	Postgres          PostgresConfig
	ExternalAPIConfig ExternalAPIConfig
}

var k = koanf.New(".")

func NewConfig() *Config {
	k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "")), "_", ".", -1)
	}), nil)

	newConfig := &Config{
		Env:  k.String("env"),
		Port: k.String("port"),
		Postgres: PostgresConfig{
			Host:     k.String("postgres.host"),
			User:     k.String("postgres.user"),
			Password: k.String("postgres.password"),
			DBName:   k.String("postgres.db"),
			Port:     k.String("postgres.port"),
		},
		ExternalAPIConfig: ExternalAPIConfig{
			FmpApiKey: k.String("fmpkey"),
		},
	}

	fmt.Printf("config loaded: %+v\n", newConfig)

	return newConfig
}
