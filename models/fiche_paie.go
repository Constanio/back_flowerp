package models

import (
	"time"
	"gorm.io/gorm"
)

type StatutPaie string

const (
	StatutBrouillon StatutPaie = "brouillon"
	StatutTraite    StatutPaie = "traite"
	StatutPaye      StatutPaie = "paye"
)

type FichePaie struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UtilisateurID uint           `gorm:"uniqueIndex:idx_user_month_year" json:"utilisateur_id"`
	Utilisateur   Utilisateur    `gorm:"foreignKey:UtilisateurID" json:"utilisateur"`
	Mois          int            `gorm:"uniqueIndex:idx_user_month_year" json:"mois"`
	Annee         int            `gorm:"uniqueIndex:idx_user_month_year" json:"annee"`
	SalaireBase   float64        `json:"salaire_base"`
	Primes        float64        `gorm:"default:0" json:"primes"`
	Deductions    float64        `gorm:"default:0" json:"deductions"`
	SalaireNet    float64        `gorm:"->;column:salaire_net" json:"salaire_net"`
	DatePaiement  *time.Time     `json:"date_paiement"`
	Statut        StatutPaie     `gorm:"default:'brouillon'" json:"statut"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
