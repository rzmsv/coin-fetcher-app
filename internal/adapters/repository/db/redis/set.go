package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (r redisDB) Set(ctx context.Context, key string, value interface{}, expires time.Duration) (err error) {
	finalValue := value

	if valueMarshaled, err1 := json.Marshal(value); err1 == nil {
		finalValue = valueMarshaled
	}
	aiwKey := keyGenerator(key)
	_, err = r.Redis.Set(ctx, aiwKey, finalValue, expires).Result()
	if err != nil {
		return err
	}
	return
}
