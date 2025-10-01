package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r redisDB) Get(key string) (res string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
