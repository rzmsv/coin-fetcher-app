package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (r redisDB) Get(ctx context.Context, key string) (res string, err error) {
	aiwKey := keyGenerator(key)
	res, err = r.Redis.Get(ctx, aiwKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", err
		}
		return
	}
	return
}
