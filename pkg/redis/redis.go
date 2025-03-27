package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Creds struct {
	Addr     string // host:port
	Password string
	DB       int
}

// NewRDS инициализирует клиента Redis
func NewRDS(creds Creds) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     creds.Addr,
		Password: creds.Password,
		DB:       creds.DB,
	})
	return rdb
}

// ParseRedisURL разбирает URL вида redis://:password@host:6379/0
func ParseRedisURL(url string) (Creds, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return Creds{}, fmt.Errorf("parse redis url error: %w", err)
	}

	return Creds{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	}, nil
}
