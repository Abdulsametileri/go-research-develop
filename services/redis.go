package services

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
	"github.com/go-redis/redis/v8"
)

// brew services restart redis
type Redis struct{}

func (t *Redis) setup() {
	redisConfig := config.Get().RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	pong, err := rdb.Ping(context.TODO()).Result()
	fmt.Println(err)
	fmt.Println("Redis PİNG: " + pong)
}
