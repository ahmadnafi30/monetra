package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type IOTPRepository interface {
	Save(ctx context.Context, email, code string, expiration time.Duration) error
	Get(ctx context.Context, email string) (string, error)
	Delete(ctx context.Context, email string) error
}

type otpRepository struct {
	redisClient *redis.Client
}

func NewOTPRepository(redisClient *redis.Client) IOTPRepository {
	return &otpRepository{
		redisClient: redisClient,
	}
}

func (r *otpRepository) Save(ctx context.Context, email, code string, expiration time.Duration) error {
	key := fmt.Sprintf("otp:%s", email)
	return r.redisClient.Set(ctx, key, code, expiration).Err()
}

func (r *otpRepository) Get(ctx context.Context, email string) (string, error) {
	key := fmt.Sprintf("otp:%s", email)
	return r.redisClient.Get(ctx, key).Result()
}

func (r *otpRepository) Delete(ctx context.Context, email string) error {
	key := fmt.Sprintf("otp:%s", email)
	return r.redisClient.Del(ctx, key).Err()
}
