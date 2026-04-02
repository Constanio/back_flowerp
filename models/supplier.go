package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null;index" json:"organization_id"`
	Name           string         `gorm:"not null" json:"name"`
	ContactName    string         `json:"contact_name"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	Address        string         `json:"address"`
	Category       string         `json:"category"` // ex: Matières premières, Services, Logistique
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *Supplier) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
