package models

import (
	"time"
	"gorm.io/gorm"
)

type EvaluationPerformance struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	UtilisateurID  uint           `json:"utilisateur_id"`
	Utilisateur    Utilisateur    `gorm:"foreignKey:UtilisateurID" json:"utilisateur"`
	EvaluateurID   uint           `json:"evaluateur_id"`
	Evaluateur     Utilisateur    `gorm:"foreignKey:EvaluateurID" json:"evaluateur"`
	PeriodeDebut   time.Time      `json:"periode_debut"`
	PeriodeFin     time.Time      `json:"periode_fin"`
	DateEvaluation time.Time      `json:"date_evaluation"`
	Score          float64        `json:"score"`
	Commentaires   string         `json:"commentaires"`
	Objectifs      string         `json:"objectifs"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
