package services

import (
	"sirh/database"
	"sirh/models"
)

// DÉPARTEMENTS
func GetAllDepartements() ([]models.Departement, error) {
	var deps []models.Departement
	err := database.DB.Preload("Manager").Find(&deps).Error
	return deps, err
}

func CreateDepartement(d *models.Departement) error {
	return database.DB.Create(d).Error
}

// POSTES
func GetAllPostes() ([]models.Poste, error) {
	var postes []models.Poste
	err := database.DB.Preload("Departement").Find(&postes).Error
	return postes, err
}

func CreatePoste(p *models.Poste) error {
	return database.DB.Create(p).Error
}
