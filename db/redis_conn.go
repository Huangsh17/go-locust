package db

import (
	"github.com/gomodule/redigo/redis"
	"os"
)

var redisPool redis.Pool

const DefaultRedisHost = "127.0.0.1:6379"

func GetRedisHost() string {
	var RedisHost = os.Getenv("REDIS_HOST")
	if len(RedisHost) != 0 {
		return DefaultRedisHost
	}

	return RedisHost
}

func RedisInit() {
	redisPool = redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", GetRedisHost())
		},
		MaxIdle:         0,
		MaxActive:       0,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func GetRedisConn() redis.Conn {
	return redisPool.Get()
}
