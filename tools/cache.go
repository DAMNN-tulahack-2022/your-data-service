package tools

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
)

var Cache *ExternalCache

type ExternalCache struct {
	log *YourLogger

	Client *redis.Client
	ctx    context.Context
}

func NewExternalCache(ctx context.Context, config *GenericEnvAppConfig, log *YourLogger) error {
	Cache = &ExternalCache{
		ctx: ctx,
		log: log,
		Client: redis.NewClient(&redis.Options{
			Addr:       config.CacheAddr,
			Password:   config.CachePassword,
			DB:         config.CacheDbCursor,
			MaxRetries: config.CacheConnectionMaxRetries,
		}),
	}

	if err := Cache.Client.Ping().Err(); err != nil {
		return err
	}

	return nil
}

func (c *ExternalCache) Set(key string, value interface{}) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.Set(key, v, 12*time.Hour).Err()
}

func (c *ExternalCache) SetWithTimeout(key string, value interface{}, ttl time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Client.Set(key, v, ttl).Err()
}

func (c *ExternalCache) Get(value interface{}, key string) error {
	data, err := c.Client.Get(key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &value)
}

func (c *ExternalCache) Keys(pattern string) []string {
	keys, err := c.Client.Keys(pattern).Result()
	if err != nil {
		keys = make([]string, 0)
	}
	return keys
}

func (c *ExternalCache) Del(key string) error {
	return c.Client.Del(key).Err()
}

func (c *ExternalCache) HGet(key, field string, value interface{}) error {
	data, err := c.Client.HGet(key, field).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &value)
}

func (c *ExternalCache) HSet(key, field string, value interface{}) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = c.Client.HSet(key, field, v).Err()
	return err
}

func (c *ExternalCache) HDel(key string, fields ...string) (result int64, err error) {
	result, err = c.Client.HDel(key, fields...).Result()
	return
}

func (c ExternalCache) Close() {
	if err := c.Client.Close(); err != nil {
		log.Error().AnErr("Error close cache connection:", err)
		return
	}
}
