package model

import (
	"time"
)

type Transaction struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id"`
	ProductID uint       `json:"product_id"` // ✅ Tambahkan ini
	Amount    float64    `json:"amount"`
	CardToken string     `json:"card_token"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
