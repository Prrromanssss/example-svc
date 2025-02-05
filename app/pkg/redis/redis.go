package redis

import (
	"context"
	"net"

	"github.com/redis/go-redis/v9"
)

type Config interface {
	GetHost() string
	GetPort() string
	GetPassword() string
	GetDB() int
	GetAddress() string
}

type ConfigModel struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func (c ConfigModel) GetHost() string {
	return c.Host
}

func (c ConfigModel) GetPort() string {
	return c.Port
}

func (c ConfigModel) GetPassword() string {
	return c.Password
}

func (c ConfigModel) GetDB() int {
	return c.DB
}

func (c ConfigModel) GetAddress() string {
	return net.JoinHostPort(c.GetHost(), c.GetPort())
}

func InitRedisDB(cfg Config) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     cfg.GetAddress(),
		Password: cfg.GetPassword(),
		DB:       cfg.GetDB(),
	}

	client := redis.NewClient(opts)
	result := client.Ping(context.Background())
	if result.Err() != nil {
		return nil, result.Err()
	}

	return client, nil
}
