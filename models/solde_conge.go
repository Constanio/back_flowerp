package models

import (
	"time"
	"gorm.io/gorm"
)

type SoldeConge struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UtilisateurID uint           `gorm:"uniqueIndex:idx_user_type_year" json:"utilisateur_id"`
	Utilisateur   Utilisateur    `gorm:"foreignKey:UtilisateurID" json:"utilisateur"`
	TypeCongeID   uint           `gorm:"uniqueIndex:idx_user_type_year" json:"type_conge_id"`
	TypeConge     TypeConge      `gorm:"foreignKey:TypeCongeID" json:"type_conge"`
	Annee         int            `gorm:"uniqueIndex:idx_user_type_year" json:"annee"`
	TotalJours    float64        `gorm:"not null;default:0" json:"total_jours"`
	JoursUtilises float64        `gorm:"default:0" json:"jours_utilises"`
	JoursRestants float64        `gorm:"->;column:jours_restants;default:0" json:"jours_restants"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
