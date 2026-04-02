package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string         `gorm:"unique;not null" json:"name"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Address   string         `json:"address"`
	VATNumber string         `json:"vat_number"`
	LogoURL   string         `json:"logo_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (o *Organization) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}
