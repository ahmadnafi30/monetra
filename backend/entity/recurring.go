package entity

import (
	"time"

	"github.com/google/uuid"
)

type RecurringTransaction struct {
	ID          uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	UserID      uuid.UUID  `json:"userId" gorm:"type:char(36);not null;index"`
	CategoryID  uuid.UUID  `json:"categoryId" gorm:"type:char(36);index"`
	Type        string     `json:"type" gorm:"type:enum('income','expense');not null"`
	Amount      float64    `json:"amount" gorm:"type:decimal(15,2);not null"`
	Description string     `json:"description" gorm:"type:text"`
	Frequency   string     `json:"frequency" gorm:"type:enum('daily','weekly','monthly','yearly');not null"`
	StartDate   time.Time  `json:"startDate" gorm:"not null"`
	EndDate     *time.Time `json:"endDate"`
	NextDate    time.Time  `json:"nextDate" gorm:"not null;index"` // next execution date
	IsActive    bool       `json:"isActive" gorm:"default:true"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relations
	User     User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}
