package models

import (
	"fmt"
	"time"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleRH      Role = "rh"
	RoleManager Role = "manager"
	RoleEmploye Role = "employe"
)

type Statut string

const (
	StatutActif   Statut = "actif"
	StatutInactif Statut = "inactif"
	StatutTermine Statut = "termine"
)

type Utilisateur struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	Matricule          string         `gorm:"unique;not null" json:"matricule"`
	Email              string         `gorm:"unique;not null" json:"email"`
	MotDePasseHash     string         `gorm:"not null" json:"-"`
	Prenom             string         `gorm:"not null" json:"prenom"`
	Nom                string         `gorm:"not null" json:"nom"`
	DateDeNaissance    *time.Time     `json:"date_de_naissance"`
	Genre              string         `json:"genre"`
	Telephone          string         `json:"telephone"`
	Adresse            string         `json:"adresse"`
	PhotoURL           string         `json:"photo_url"`
	Role               Role           `gorm:"type:varchar(20);not null;default:'employe'" json:"role"`
	Statut             Statut         `gorm:"type:varchar(20);default:'actif'" json:"statut"`
	DateEmbauche       time.Time      `gorm:"not null" json:"date_embauche"`
	DateFinContrat     *time.Time     `json:"date_fin_contrat"`
	DepartementID      *uint          `json:"departement_id"`
	Departement        *Departement   `gorm:"foreignKey:DepartementID" json:"departement"`
	PosteID            *uint          `json:"poste_id"`
	Poste              *Poste         `gorm:"foreignKey:PosteID" json:"poste"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *Utilisateur) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Matricule == "" {
		var count int64
		tx.Model(&Utilisateur{}).Count(&count)
		u.Matricule = fmt.Sprintf("EMP%04d", count+1)
	}
	return
}

func (u *Utilisateur) AfterCreate(tx *gorm.DB) (err error) {
	var types []TypeConge
	tx.Find(&types)

	for _, t := range types {
		solde := SoldeConge{
			UtilisateurID: u.ID,
			TypeCongeID:   t.ID,
			Annee:         time.Now().Year(),
			TotalJours:    float64(t.JoursParAn),
			JoursUtilises: 0,
		}
		tx.Create(&solde)
	}
	return
}
