package handlers

import (
	"sirh/models"
	"sirh/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DÉPARTEMENTS
func GetDepartements(c *gin.Context) {
	deps, err := services.GetAllDepartements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deps)
}

func CreateDepartement(c *gin.Context) {
	var input models.Departement
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateDepartement(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// POSTES
func GetPostes(c *gin.Context) {
	postes, err := services.GetAllPostes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, postes)
}

func CreatePoste(c *gin.Context) {
	var input models.Poste
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreatePoste(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}
