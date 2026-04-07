package services

import (
	"sirh/database"
	"sirh/models"
)

// TYPES DE CONGÉS
func GetAllTypesConges() ([]models.TypeConge, error) {
	var types []models.TypeConge
	err := database.DB.Find(&types).Error
	return types, err
}

// DEMANDES DE CONGÉS
func CreateDemandeConge(d *models.DemandeConge) error {
	return database.DB.Create(d).Error
}

func GetDemandesByUser(userID uint) ([]models.DemandeConge, error) {
	var demandes []models.DemandeConge
	err := database.DB.Preload("TypeConge").Where("utilisateur_id = ?", userID).Find(&demandes).Error
	return demandes, err
}

func UpdateStatutDemande(id uint, statut models.StatutDemande, approuvePar uint) error {
	return database.DB.Model(&models.DemandeConge{}).Where("id = ?", id).Updates(map[string]interface{}{
		"statut": statut,
		"approuve_par_id": approuvePar,
	}).Error
}

// SOLDES
func GetSoldesByUser(userID uint) ([]models.SoldeConge, error) {
	var soldes []models.SoldeConge
	err := database.DB.Preload("TypeConge").Where("utilisateur_id = ?", userID).Find(&soldes).Error
	return soldes, err
}
