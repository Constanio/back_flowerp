package handlers

import (
	"sirh/database"
	"sirh/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardStats struct {
	UtilisateurCount int64 `json:"utilisateur_count"`
	DemandesEnAttente int64 `json:"demandes_en_attente"`
	DepartementCount int64 `json:"departement_count"`
	PosteCount       int64 `json:"poste_count"`
}

func GetDashboardStats(c *gin.Context) {
	var stats DashboardStats

	database.DB.Model(&models.Utilisateur{}).Count(&stats.UtilisateurCount)
	database.DB.Model(&models.DemandeConge{}).Where("statut = ?", "en_attente").Count(&stats.DemandesEnAttente)
	database.DB.Model(&models.Departement{}).Count(&stats.DepartementCount)
	database.DB.Model(&models.Poste{}).Count(&stats.PosteCount)

	c.JSON(http.StatusOK, stats)
}

func GetMonthlyRevenue(c *gin.Context) {
	// Cette fonction pourrait être supprimée ou transformée en statistiques RH (ex: embauches par mois)
	c.JSON(http.StatusOK, gin.H{"message": "Non implémenté pour le module RH pur"})
}
