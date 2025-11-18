package repository

import (
	"github.com/ahmadnafi30/monetra/backend/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *entity.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}
