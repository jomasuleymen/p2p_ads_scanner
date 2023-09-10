package initializers

import (
	"net"
	"os"

	"github.com/redis/go-redis/v9"
)

var cache *redis.Client

func GetRedisDB() *redis.Client {

	if cache != nil {
		return cache
	}

	// dbPort := os.Getenv("REDIS_PORT")
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cache = rdb

	return rdb
}
