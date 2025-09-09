package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID         uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;not null"`
	Title      string    `json:"title" gorm:"type:varchar(255);not null"`
	AuthorID   uuid.UUID `json:"author_id" gorm:"type:varchar(36);not null"`
	CategoryID uuid.UUID `json:"category_id" gorm:"type:varchar(36);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (b Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
