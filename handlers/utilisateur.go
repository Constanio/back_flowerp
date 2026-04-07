package handlers

import (
	"sirh/models"
	"sirh/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUtilisateurs(c *gin.Context) {
	utils, err := services.GetAllUtilisateurs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer les utilisateurs"})
		return
	}
	c.JSON(http.StatusOK, utils)
}

func GetUtilisateur(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	util, err := services.GetUtilisateurByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}
	c.JSON(http.StatusOK, util)
}

func CreateUtilisateur(c *gin.Context) {
	var input models.Utilisateur
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateUtilisateur(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func UpdateUtilisateur(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	util, err := services.GetUtilisateurByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Utilisateur non trouvé"})
		return
	}

	var input struct {
		Prenom        string `json:"prenom"`
		Nom           string `json:"nom"`
		Email         string `json:"email"`
		Role          string `json:"role"`
		DepartementID *uint  `json:"departement_id"`
		PosteID       *uint  `json:"poste_id"`
		Statut        string `json:"statut"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mise à jour des champs
	util.Prenom = input.Prenom
	util.Nom = input.Nom
	util.Email = input.Email
	util.Role = models.Role(input.Role)
	util.DepartementID = input.DepartementID
	util.PosteID = input.PosteID
	if input.Statut != "" {
		util.Statut = models.Statut(input.Statut)
	}

	if err := services.UpdateUtilisateur(&util); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour"})
		return
	}
	c.JSON(http.StatusOK, util)
}

func DeleteUtilisateur(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteUtilisateur(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Utilisateur supprimé"})
}
