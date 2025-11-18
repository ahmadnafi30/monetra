package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository struct {
	User        IUserRepository
	OTP         IOTPRepository
	Category    CategoryRepository
	Transaction TransactionRepository
}

func NewRepository(db *gorm.DB, redisClient *redis.Client) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		OTP:         NewOTPRepository(redisClient),
		Category:    NewCategoryRepository(db),
		Transaction: NewTransactionRepository(db),
	}
}
