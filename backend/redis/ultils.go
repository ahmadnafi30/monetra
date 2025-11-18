package redis

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	Password: "",
	DB:       0,
})

func SetValue(key string, value string) error {
	return redisClient.Set(ctx, key, value, 1*time.Minute).Err()
}

func GetValue(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}
