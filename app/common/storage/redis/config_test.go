package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisConfig_URL(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		expected string
	}{
		{
			name: "with password",
			config: Config{
				Address:  "a.b.c",
				Port:     "6379",
				Password: "mypass",
			},
			expected: "redis://:mypass@a.b.c:6379",
		},
		{
			name: "with password and tls",
			config: Config{
				Address:  "a.b.c",
				Port:     "6379",
				Password: "mypass",
				UseTLS:   true,
			},
			expected: "rediss://:mypass@a.b.c:6379",
		},
		{
			name: "with tls",
			config: Config{
				Address: "a.b.c",
				Port:    "6379",
				UseTLS:  true,
			},
			expected: "rediss://a.b.c:6379",
		},
		{
			name: "without tls and password",
			config: Config{
				Address: "a.b.c",
				Port:    "6379",
			},
			expected: "redis://a.b.c:6379",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.config.URL())
		})
	}
}
