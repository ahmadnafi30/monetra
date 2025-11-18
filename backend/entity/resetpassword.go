package entity

import (
	"time"

	"github.com/google/uuid"
)

type PasswordReset struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Otp       string    `json:"otp" gorm:"type:varchar(6);not null"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
