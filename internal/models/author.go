package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Books     []Book    `json:"-" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == (uuid.UUID{}) {
		a.ID = uuid.New()
	}
	return
}
