package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null;index" json:"organization_id"`
	Name           string         `gorm:"not null" json:"name"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	Address        string         `json:"address"`
	Status         string         `gorm:"default:prospect" json:"status"` // prospect, active, inactive
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
