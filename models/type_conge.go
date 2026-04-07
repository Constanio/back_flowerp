package models

import (
	"time"
	"gorm.io/gorm"
)

type TypeConge struct {
	ID                    uint           `gorm:"primaryKey" json:"id"`
	Nom                   string         `gorm:"not null" json:"nom"`
	Description           string         `json:"description"`
	JoursParAn            int            `gorm:"default:0" json:"jours_par_an"`
	NecessiteApprobation  bool           `gorm:"default:true" json:"necessite_approbation"`
	Couleur               string         `json:"couleur"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`
}
