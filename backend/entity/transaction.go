package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	UserID      uuid.UUID  `json:"userId" gorm:"type:char(36);not null;index"`
	CategoryID  uuid.UUID  `json:"categoryId" gorm:"type:char(36);index"`
	Type        string     `json:"type" gorm:"type:enum('income','expense');not null"`
	Amount      float64    `json:"amount" gorm:"type:decimal(15,2);not null"`
	Description string     `json:"description" gorm:"type:text"`
	Date        time.Time  `json:"date" gorm:"not null;index"` // transaction date
	IsRecurring bool       `json:"isRecurring" gorm:"default:false"`
	InvoiceURL  string     `json:"invoiceUrl" gorm:"type:text"` // URL to invoice image
	Notes       string     `json:"notes" gorm:"type:text"`
	Tags        string     `json:"tags" gorm:"type:varchar(255)"` // comma-separated tags
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" gorm:"index"` // soft delete

	// Relations
	User     User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}
