package redisdb

import (
	"github.com/go-redis/redis"
)

// RedisConnector stores all the config details for redis server
type RedisConnector struct {
	Address        string   `toml:"address"`
}

// RedisDB stores the config struct and connection object
type RedisDB struct {
	Config RedisConnector `toml:"RedisConfig"`
	Client *redis.Client
}

