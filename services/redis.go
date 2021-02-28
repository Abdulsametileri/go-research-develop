package services

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
	"github.com/go-redis/redis/v8"
	"log"
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
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		log.Panic("Redise PİNG atılamadı." + err.Error())
	}

	fmt.Println("REDİS AYAĞA KALKTI")
}
