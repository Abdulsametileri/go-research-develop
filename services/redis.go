package services

import (
	"context"
	"fmt"
	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// brew services restart redis
type Redis struct{}

func (t *Redis) setup() {
	redisConfig := config.Get().RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Username: redisConfig.Username,
		Password: redisConfig.Password,
	})

	errf := fmt.Sprintf("%s %s", redisConfig.Password, redisConfig.Host)
	fmt.Println(errf)
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		log.Panic("Redise PİNG atılamadı." + err.Error())
	}

	fmt.Println("REDİS AYAĞA KALKTI")

	cvp, errSet := rdb.Set(context.TODO(), "bismillah", 3, time.Minute*60).Result()
	fmt.Println(cvp, errSet)
}
