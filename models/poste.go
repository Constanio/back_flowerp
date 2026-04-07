package models

import (
	"time"
	"gorm.io/gorm"
)

type Poste struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Titre         string         `gorm:"not null" json:"titre"`
	DepartementID *uint          `json:"departement_id"`
	Departement   *Departement    `gorm:"foreignKey:DepartementID" json:"departement"`
	Description   string         `json:"description"`
	SalaireMin    float64        `json:"salaire_min"`
	SalaireMax    float64        `json:"salaire_max"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
