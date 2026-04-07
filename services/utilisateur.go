package services

import (
	"sirh/database"
	"sirh/models"
)

func GetAllUtilisateurs() ([]models.Utilisateur, error) {
	var utilisateurs []models.Utilisateur
	err := database.DB.Preload("Departement").Preload("Poste").Find(&utilisateurs).Error
	return utilisateurs, err
}

func GetUtilisateurByID(id uint) (models.Utilisateur, error) {
	var utilisateur models.Utilisateur
	err := database.DB.Preload("Departement").Preload("Poste").First(&utilisateur, id).Error
	return utilisateur, err
}

func CreateUtilisateur(u *models.Utilisateur) error {
	return database.DB.Create(u).Error
}

func UpdateUtilisateur(u *models.Utilisateur) error {
	return database.DB.Save(u).Error
}

func DeleteUtilisateur(id uint) error {
	return database.DB.Delete(&models.Utilisateur{}, id).Error
}
