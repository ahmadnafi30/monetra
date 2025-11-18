package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Provider  string    `json:"provider" gorm:"type:varchar(50);default:'manual'"`
	Avatar    string    `json:"avatar" gorm:"type:text"`
	Phone     string    `json:"phone" gorm:"type:varchar(20)"`
	Currency  string    `json:"currency" gorm:"type:varchar(10);default:'IDR'"` // Default currency
	Timezone  string    `json:"timezone" gorm:"type:varchar(50);default:'Asia/Jakarta'"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	// Relations
	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:UserID"`
	Categories   []Category    `json:"categories,omitempty" gorm:"foreignKey:UserID"`
	Budgets      []Budget      `json:"budgets,omitempty" gorm:"foreignKey:UserID"`
}
