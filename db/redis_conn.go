package db

import (
	"github.com/gomodule/redigo/redis"
	"go-locust/config"
)

var redisPool redis.Pool

func RedisInit() {
	redisPool = redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", config.HOST_REDIS)
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
