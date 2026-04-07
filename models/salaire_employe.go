package models

import (
	"time"
	"gorm.io/gorm"
)

type SalaireEmploye struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UtilisateurID uint           `json:"utilisateur_id"`
	Utilisateur   Utilisateur    `gorm:"foreignKey:UtilisateurID" json:"utilisateur"`
	SalaireBase   float64        `gorm:"not null" json:"salaire_base"`
	DateDebut     time.Time      `gorm:"not null" json:"date_debut"`
	DateFin       *time.Time     `json:"date_fin"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
