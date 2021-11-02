package redis

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Address            string        `envconfig:"REDIS_ADDR" default:"0.0.0.0"`
	Port               string        `envconfig:"REDIS_PORT" default:"6379"`
	Password           string        `envconfig:"REDIS_PASSWORD"`
	UseTLS             bool          `envconfig:"REDIS_USE_TLS" default:"false"`
	MaxIdle            int           `envconfig:"REDIS_MAX_IDLE" default:"100"`
	MaxActive          int           `envconfig:"REDIS_MAX_ACTIVE" default:"1000"`
	IdleTimeout        time.Duration `envconfig:"REDIS_IDLE_TIMEOUT" default:"1m"`
	DialConnectTimeout time.Duration `envconfig:"REDIS_CONNECT_TIMEOUT" default:"1s"`
	DialReadTimeout    time.Duration `envconfig:"REDIS_READ_TIMEOUT" default:"300ms"`
	DialWriteTimeout   time.Duration `envconfig:"REDIS_WRITE_TIMEOUT" default:"300ms"`
}

func (c Config) Addr() string {
	return strings.Join([]string{c.Address, c.Port}, ":")
}

func (c Config) URL() string {
	scheme := "redis"
	if c.UseTLS {
		scheme = "rediss"
	}
	addr := c.Addr()
	if c.Password != "" {
		addr = fmt.Sprintf(":%s@%s", c.Password, addr)
	}
	return fmt.Sprintf("%s://%s", scheme, addr)
}
