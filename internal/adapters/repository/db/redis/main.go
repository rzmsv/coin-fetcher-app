package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/username/coin-fetcher-app/config"
	"github.com/username/coin-fetcher-app/internal/domain/interfaces"
)

type redisDB struct {
	Redis *redis.Client
}

// func NewRedis(host, user, password string, db, port int) *redisDB {
func NewRedis(config *config.AppConfig) interfaces.Redis {
	db, _ := strconv.Atoi(config.Configs("REDIS_DB"))
	port, _ := strconv.Atoi(config.Configs("REDIS_PORT"))

	addr := fmt.Sprintf("%s:%d", config.Configs("REDIS_HOST"), port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: config.Configs("REDIS_USER"),
		Password: config.Configs("REDIS_PASSWORD"),
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
