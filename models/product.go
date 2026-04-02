package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null;index" json:"organization_id"`
	Name           string         `gorm:"not null" json:"name"`
	Description    string         `json:"description"`
	Category       string         `json:"category"`
	Price          float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock          int            `gorm:"default:0" json:"stock"`
	MinStock       int            `gorm:"default:5" json:"min_stock"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
