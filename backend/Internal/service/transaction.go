package service

import (
	"github.com/ahmadnafi30/monetra/backend/Internal/repository"
	"github.com/ahmadnafi30/monetra/backend/model"
	"github.com/ahmadnafi30/monetra/backend/pkg/response"
	"github.com/google/uuid"
)

type ITransactionService interface {
}

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) ITransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(transaction *model.CreateTransactionRequest) error {
	if transaction.Amount <= 0 {
		return response.Error(400, "Amount must be greater than zero", nil)
	}

	if transaction.CategoryID == uuid.Nil {
		return response.Error(400, "Category ID is required", nil)
	}

	return nil
}

// type CreateTransactionRequest struct {
// 	CategoryID  uuid.UUID `json:"categoryId" binding:"required"`
// 	Type        string    `json:"type" binding:"required,oneof=income expense"`
// 	Amount      float64   `json:"amount" binding:"required,gt=0"`
// 	Description string    `json:"description"`
// 	Date        time.Time `json:"date" binding:"required"`
// 	IsRecurring bool      `json:"isRecurring"`
// 	Notes       string    `json:"notes"`
// 	Tags        string    `json:"tags"`
// }
