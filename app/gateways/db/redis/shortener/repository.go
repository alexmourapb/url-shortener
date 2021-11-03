package shortener

import (
	"github.com/gomodule/redigo/redis"

	"github.com/alexmourapb/url-shortener/app/domain/shortener"
)

const (
	argSet        = "SET" // Redis arg SET
	argGet        = "GET" // Redis arg SET
	argDel        = "DEL" // Redis arg SET
	argNx         = "NX"  // Redis arg set if the key does not exist
	argEx         = "EX"  // Redis arg ttl key
	ExpireKeyTime = 1800  // 30 minutes in seconds
)

var _ shortener.CacheRepository = &Repository{}

type Repository struct {
	redisPool *redis.Pool
}

func NewCacheRepository(redis *redis.Pool) *Repository {
	return &Repository{
		redisPool: redis,
	}
}
