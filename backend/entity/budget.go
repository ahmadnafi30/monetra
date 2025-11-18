package entity

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID     uuid.UUID `json:"userId" gorm:"type:char(36);not null;index"`
	CategoryID uuid.UUID `json:"categoryId" gorm:"type:char(36);index"`
	Amount     float64   `json:"amount" gorm:"type:decimal(15,2);not null"` // budget limit
	Period     string    `json:"period" gorm:"type:enum('monthly','yearly');default:'monthly'"`
	Month      int       `json:"month" gorm:"type:int"` // 1-12 for monthly
	Year       int       `json:"year" gorm:"type:int;not null"`
	Alert      bool      `json:"alert" gorm:"default:true"`                     // send alert when exceeding
	Threshold  float64   `json:"threshold" gorm:"type:decimal(5,2);default:80"` // alert at 80%
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relations
	User     User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}
