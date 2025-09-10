package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;not null"`
	Amount    float64   `json:"amount" gorm:"not null"`
	Book      Book      `json:"book" gorm:"foreignKey:ID"`
	Method    string    `json:"method" gorm:"type:varchar(50);not null"`
	Currency  string    `json:"currency" gorm:"type:varchar(10);not null"`
	Status    string    `json:"status" gorm:"type:varchar(50);not null"`
	OrderID   uuid.UUID `json:"order_id" gorm:"type:varchar(36);not null"`
	CreatedAt int64     `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64     `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == (uuid.UUID{}) {
		p.ID = uuid.New()
	}
	return
}
