package redisx

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Options struct {
	opts redis.Options
}

var defaultOptions = Options{
	opts: redis.Options{
		Addr:         ":6379",
		Username:     "",
		Password:     "",
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     64,
		MinIdleConns: 10,
	},
}

func WithAuthenticate(username, password string) func(*Options) {
	return func(options *Options) {
		options.opts.Username = username
		options.opts.Password = password
	}
}

func WithDBIndex(index int) func(*Options) {
	return func(options *Options) {
		options.opts.DB = index
	}
}

func WithTimeouts(dial, read, write time.Duration) func(*Options) {
	return func(options *Options) {
		options.opts.DialTimeout = dial
		options.opts.ReadTimeout = read
		options.opts.WriteTimeout = write
	}
}

func WithPoolConfig(size, minIdle, maxIdle int) func(*Options) {
	return func(options *Options) {
		options.opts.PoolSize = size
		options.opts.MinIdleConns = minIdle
		options.opts.MaxIdleConns = maxIdle
	}
}
