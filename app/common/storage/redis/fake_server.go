package redis

import (
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
)

func NewFakeRedisServer() *redis.Pool {
	redisServer, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	redisConn, err := redis.Dial("tcp", redisServer.Addr())
	if err != nil {
		panic(err)
	}

	return &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 1 * time.Minute,
		Dial:        func() (redis.Conn, error) { return redisConn, nil },
	}
}
