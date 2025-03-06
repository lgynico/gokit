package redisx

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	client redis.Client
}

func NewClient(addr string, opts ...func(*Options)) *Client {
	options := defaultOptions
	options.opts.Addr = addr

	for _, opt := range opts {
		opt(&options)
	}

	return &Client{
		client: *redis.NewClient(&options.opts),
	}
}

func (p *Client) Ping() error {
	return p.client.Ping(context.TODO()).Err()
}

func (p *Client) Redis() *redis.Client {
	return &p.client
}

func (p *Client) Set(key string, value any) (string, error) {
	return p.client.Set(context.TODO(), key, value, 0).Result()
}

func (p *Client) SetEX(key string, value any, expiration time.Duration) (string, error) {
	return p.client.Set(context.TODO(), key, value, expiration).Result()
}

func (p *Client) SetNX(key string, value any, expiration time.Duration) (bool, error) {
	return p.client.SetNX(context.TODO(), key, value, expiration).Result()
}

func (p *Client) Get(key string) (string, error) {
	return p.client.Get(context.TODO(), key).Result()
}

func (p *Client) Exists(key string) (bool, error) {
	result, err := p.client.Exists(context.TODO(), key).Result()
	if err != nil {
		return false, err
	}

	return result == 1, nil
}

func (p *Client) HMGet(key string, fields ...string) (map[string]any, error) {
	result, err := p.client.HMGet(context.TODO(), key, fields...).Result()
	if err != nil {
		return nil, err
	}

	m := make(map[string]any, len(result))
	for i := 0; i < len(fields); i++ {
		if result[i] == nil {
			continue
		}
		m[fields[i]] = result[i]
	}
	return m, nil
}

func (p *Client) HGet(key string) (map[string]any, error) {
	result, err := p.client.HGetAll(context.TODO(), key).Result()
	if err != nil {
		return nil, err
	}

	m := make(map[string]any, len(result))
	for k, v := range result {
		m[k] = v
	}

	return m, nil
}

func (p *Client) HMSet(key string, fields map[string]any) (int64, error) {
	return p.client.HSet(context.TODO(), key, fields).Result()
}

func (p *Client) Expire(key string, expiration time.Duration) (bool, error) {
	return p.client.Expire(context.TODO(), key, expiration).Result()
}

func (p *Client) Deadline(key string, deadline time.Time) (bool, error) {
	return p.client.ExpireAt(context.TODO(), key, deadline).Result()
}

func (p *Client) Incr(key string) (int64, error) {
	return p.client.Incr(context.TODO(), key).Result()
}

func (p *Client) Decr(key string) (int64, error) {
	return p.client.Decr(context.TODO(), key).Result()
}

func (p *Client) Del(key string) (bool, error) {
	result, err := p.client.Del(context.TODO(), key).Result()
	if err != nil {
		return false, err
	}

	return result == 1, nil
}

func (p *Client) ZAdd(key string, score float64, value any) (int64, error) {
	return p.client.ZAdd(context.TODO(), key, redis.Z{Score: score, Member: value}).Result()
}

func (p *Client) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return p.client.ZRevRangeWithScores(context.TODO(), key, start, stop).Result()
}

func (p *Client) Eval(script string, keys []string, args ...any) (any, error) {
	return p.client.Eval(context.TODO(), script, keys, args...).Result()
}

func (p *Client) EvalSha(sha string, keys []string, args ...any) (any, error) {
	return p.client.EvalSha(context.TODO(), sha, keys, args...).Result()
}
