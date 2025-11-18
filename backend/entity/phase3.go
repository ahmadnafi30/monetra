package entity

import (
	"time"

	"github.com/google/uuid"
)

type SavingsGoal struct {
	ID            uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID        uuid.UUID `json:"userId" gorm:"type:char(36);not null;index"`
	Name          string    `json:"name" gorm:"type:varchar(255);not null"`
	Description   string    `json:"description" gorm:"type:text"`
	TargetAmount  float64   `json:"targetAmount" gorm:"type:decimal(15,2);not null"`
	CurrentAmount float64   `json:"currentAmount" gorm:"type:decimal(15,2);default:0"`
	Deadline      time.Time `json:"deadline"`
	Icon          string    `json:"icon" gorm:"type:varchar(50)"`
	Color         string    `json:"color" gorm:"type:varchar(20)"`
	Status        string    `json:"status" gorm:"type:enum('active','completed','cancelled');default:'active'"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type Notification struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	UserID    uuid.UUID  `json:"userId" gorm:"type:char(36);not null;index"`
	Type      string     `json:"type" gorm:"type:varchar(50);not null"` // budget_alert, bill_reminder, etc
	Title     string     `json:"title" gorm:"type:varchar(255);not null"`
	Message   string     `json:"message" gorm:"type:text"`
	IsRead    bool       `json:"isRead" gorm:"default:false"`
	ReadAt    *time.Time `json:"readAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`

	// Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type AIInsight struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	UserID    uuid.UUID  `json:"userId" gorm:"type:char(36);not null;index"`
	Type      string     `json:"type" gorm:"type:varchar(50);not null"` // pattern, anomaly, recommendation
	Title     string     `json:"title" gorm:"type:varchar(255);not null"`
	Message   string     `json:"message" gorm:"type:text"`
	Data      string     `json:"data" gorm:"type:json"` // JSON data for visualization
	IsViewed  bool       `json:"isViewed" gorm:"default:false"`
	ViewedAt  *time.Time `json:"viewedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`

	// Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

type AuditLog struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID     uuid.UUID `json:"userId" gorm:"type:char(36);index"`
	Action     string    `json:"action" gorm:"type:varchar(100);not null"` // login, create_transaction, etc
	EntityType string    `json:"entityType" gorm:"type:varchar(50)"`       // transaction, budget, etc
	EntityID   uuid.UUID `json:"entityId" gorm:"type:char(36)"`
	IPAddress  string    `json:"ipAddress" gorm:"type:varchar(45)"`
	UserAgent  string    `json:"userAgent" gorm:"type:text"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
