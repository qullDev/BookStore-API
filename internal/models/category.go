package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;not null"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Books     []Book    `json:"-" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == (uuid.UUID{}) {
		c.ID = uuid.New()
	}
	return
}
