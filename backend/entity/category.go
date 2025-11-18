package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID    uuid.UUID `json:"userId" gorm:"type:char(36);not null;index"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Type      string    `json:"type" gorm:"type:varchar(20);not null"` // "income" or "expense"
	Icon      string    `json:"icon" gorm:"type:varchar(50)"`
	Color     string    `json:"color" gorm:"type:varchar(20)"`
	IsDefault bool      `json:"isDefault" gorm:"default:false"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	User         User          `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:CategoryID"`
}
