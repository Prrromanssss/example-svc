package config

import (
	"example-svc/pkg/postgres"
	"example-svc/pkg/redis"
)

type Config struct {
	postgresConfig postgres.ConfigModel
	redisConfig    redis.ConfigModel
}

// The data flow of configs from connectors/adapters into this package, not the other way around
