package entity

import (
	"time"

	"github.com/google/uuid"
)

type SplitBill struct {
	ID          uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	CreatorID   uuid.UUID `json:"creatorId" gorm:"type:char(36);not null;index"` // who created
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	TotalAmount float64   `json:"totalAmount" gorm:"type:decimal(15,2);not null"`
	InvoiceURL  string    `json:"invoiceUrl" gorm:"type:text"` // scanned invoice
	Date        time.Time `json:"date" gorm:"not null"`
	Status      string    `json:"status" gorm:"type:enum('pending','partial','completed');default:'pending'"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relations
	Creator      User              `json:"creator,omitempty" gorm:"foreignKey:CreatorID"`
	Participants []SplitBillMember `json:"participants,omitempty" gorm:"foreignKey:SplitBillID"`
}
