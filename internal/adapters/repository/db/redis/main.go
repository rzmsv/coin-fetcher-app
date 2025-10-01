package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type redisDB struct {
	Redis *redis.Client
}

func NewRedis(host, user, password string, db, port int) *redisDB {

	addr := fmt.Sprintf("%s:%d", host, port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: user,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("connection to the redis database couldnt be established!")
	}
	return &redisDB{
		Redis: client,
	}
}
