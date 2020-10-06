package redisdb

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)


var rDBStatus bool = false
var rDB *RedisDB

var ctx = context.Background()

// RedisInit is used to initialise redis connection and return client
func RedisInit() (*RedisDB, error) {
	rDB := &RedisDB{}
	rDB.initConfig()
	var client *redis.Client
	var err error
	client = redis.NewClient(&redis.Options{
        Addr: rDB.Config.Address,
        })
	rDB.Client = client
	pong, err := rDB.Client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Failed to ping redis server:", err)
		return nil, err
	}
	fmt.Println("Successfully initialised redis server. Ping result:", pong)
	return rDB, nil

	
}

// Get method is used to get redis client instance
func Get() *RedisDB {
	if rDBStatus == false {
		var err error
		rDB, err = RedisInit()
		if err != nil {
			fmt.Println("Unable to get redis client instance:", err)
		} else {
			fmt.Println("Successfully initialised redis client")
			rDBStatus = true
		}
	}
	return rDB
}


