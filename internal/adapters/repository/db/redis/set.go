package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (r redisDB) Set(key string, value interface{}, expires time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
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
