package model

import (
	"time"

	"github.com/google/uuid"
)

// Category Models
type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Type  string `json:"type" binding:"required,oneof=income expense"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

type UpdateCategoryRequest struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

type CategoryParam struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Type   string // income or expense
}

// Transaction Models
type CreateTransactionRequest struct {
	CategoryID  uuid.UUID `json:"categoryId" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=income expense"`
	Amount      float64   `json:"amount" binding:"required,gt=0"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
	IsRecurring bool      `json:"isRecurring"`
	Notes       string    `json:"notes"`
	Tags        string    `json:"tags"`
}

type UpdateTransactionRequest struct {
	CategoryID  uuid.UUID `json:"categoryId"`
	Amount      float64   `json:"amount" binding:"gt=0"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Notes       string    `json:"notes"`
	Tags        string    `json:"tags"`
}

type TransactionParam struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CategoryID uuid.UUID
	Type       string
	StartDate  time.Time
	EndDate    time.Time
	Page       int
	Limit      int
}

type TransactionListResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
	Total        int64                 `json:"total"`
	Page         int                   `json:"page"`
	Limit        int                   `json:"limit"`
}

type TransactionResponse struct {
	ID          uuid.UUID        `json:"id"`
	CategoryID  uuid.UUID        `json:"categoryId"`
	Category    CategoryResponse `json:"category"`
	Type        string           `json:"type"`
	Amount      float64          `json:"amount"`
	Description string           `json:"description"`
	Date        time.Time        `json:"date"`
	IsRecurring bool             `json:"isRecurring"`
	InvoiceURL  string           `json:"invoiceUrl"`
	Notes       string           `json:"notes"`
	Tags        string           `json:"tags"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type CategoryResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Type  string    `json:"type"`
	Icon  string    `json:"icon"`
	Color string    `json:"color"`
}

// Budget Models
type CreateBudgetRequest struct {
	CategoryID uuid.UUID `json:"categoryId" binding:"required"`
	Amount     float64   `json:"amount" binding:"required,gt=0"`
	Period     string    `json:"period" binding:"required,oneof=monthly yearly"`
	Month      int       `json:"month" binding:"min=1,max=12"`
	Year       int       `json:"year" binding:"required"`
	Alert      bool      `json:"alert"`
	Threshold  float64   `json:"threshold" binding:"min=0,max=100"`
}

type UpdateBudgetRequest struct {
	Amount    float64 `json:"amount" binding:"gt=0"`
	Alert     bool    `json:"alert"`
	Threshold float64 `json:"threshold" binding:"min=0,max=100"`
}

type BudgetParam struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CategoryID uuid.UUID
	Month      int
	Year       int
	Period     string
}

type BudgetResponse struct {
	ID         uuid.UUID        `json:"id"`
	CategoryID uuid.UUID        `json:"categoryId"`
	Category   CategoryResponse `json:"category"`
	Amount     float64          `json:"amount"`
	Spent      float64          `json:"spent"`
	Remaining  float64          `json:"remaining"`
	Percentage float64          `json:"percentage"`
	Period     string           `json:"period"`
	Month      int              `json:"month"`
	Year       int              `json:"year"`
	Alert      bool             `json:"alert"`
	Threshold  float64          `json:"threshold"`
	Status     string           `json:"status"` // safe, warning, exceeded
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt"`
}
