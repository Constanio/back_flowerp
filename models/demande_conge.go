package models

import (
	"time"
	"gorm.io/gorm"
)

type StatutDemande string

const (
	StatutEnAttente StatutDemande = "en_attente"
	StatutApprouve  StatutDemande = "approuve"
	StatutRefuse    StatutDemande = "refuse"
	StatutAnnule    StatutDemande = "annule"
)

type DemandeConge struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	UtilisateurID    uint           `gorm:"not null" json:"utilisateur_id"`
	Utilisateur      Utilisateur    `gorm:"foreignKey:UtilisateurID" json:"utilisateur"`
	TypeCongeID      uint           `gorm:"not null" json:"type_conge_id"`
	TypeConge        TypeConge      `gorm:"foreignKey:TypeCongeID" json:"type_conge"`
	DateDebut        time.Time      `gorm:"not null" json:"date_debut"`
	DateFin          time.Time      `gorm:"not null" json:"date_fin"`
	JoursDemandes    float64        `gorm:"not null" json:"jours_demandes"`
	Motif            string         `json:"motif"`
	Statut           StatutDemande  `gorm:"default:'en_attente'" json:"statut"`
	ApprouveParID    *uint          `json:"approuve_par_id"`
	ApprouvePar      *Utilisateur   `gorm:"foreignKey:ApprouveParID" json:"approuve_par"`
	DateApprobation  *time.Time     `json:"date_approbation"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
