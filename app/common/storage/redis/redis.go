package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func InitPool(cfg Config) (*redis.Pool, error) {
	redisPool := &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", cfg.Addr(),
				redis.DialPassword(cfg.Password),
				redis.DialUseTLS(cfg.UseTLS),
				redis.DialConnectTimeout(cfg.DialConnectTimeout),
				redis.DialReadTimeout(cfg.DialReadTimeout),
				redis.DialWriteTimeout(cfg.DialWriteTimeout))
			if err != nil {
				return nil, fmt.Errorf(`could not connect to redis: %w`, err)
			}
			return conn, nil
		},
	}

	err := Ping(redisPool)
	if err != nil {
		return nil, err
	}

	return redisPool, nil
}

func Ping(pool *redis.Pool) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return err
	}

	return nil
}
