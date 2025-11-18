package entity

import (
	"time"

	"github.com/google/uuid"
)

type SplitBillMember struct {
	ID          uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	SplitBillID uuid.UUID  `json:"splitBillId" gorm:"type:char(36);not null;index"`
	UserID      *uuid.UUID `json:"userId" gorm:"type:char(36);index"`      // nullable for non-users
	Name        string     `json:"name" gorm:"type:varchar(255);not null"` // for non-registered users
	Email       string     `json:"email" gorm:"type:varchar(255)"`
	Phone       string     `json:"phone" gorm:"type:varchar(20)"`
	Amount      float64    `json:"amount" gorm:"type:decimal(15,2);not null"` // their share
	IsPaid      bool       `json:"isPaid" gorm:"default:false"`
	PaidAt      *time.Time `json:"paidAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relations
	SplitBill SplitBill `json:"splitBill,omitempty" gorm:"foreignKey:SplitBillID"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
