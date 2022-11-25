package tools

import (
	"context"

	"github.com/go-redis/redis"
)

type ExternalCache struct {
	log *YourLogger

	client *redis.Client
	ctx    context.Context
}

func NewExternalCache(ctx context.Context, config *GenericEnvAppConfig, log *YourLogger) (cache *ExternalCache, err error) {
	cache = &ExternalCache{
		ctx: ctx,
		log: log,
		client: redis.NewClient(&redis.Options{
			Addr:       config.CacheHost,
			Password:   config.CachePassword,
			DB:         config.CacheDbCursor,
			MaxRetries: config.CacheConnectionMaxRetries,
		}),
	}

	if err = cache.client.Ping().Err(); err != nil {
		return 
	}

	return
}
