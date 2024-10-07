package controllers

import (
	"net/http"

	"simcard/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewDB struct {
	DB *gorm.DB
}

// CreateSim is function which used to Create the sim

func (views *NewDB) CreateSim(c *gin.Context) {
	var sim models.SimcardSchema

	if err := c.ShouldBindJSON(&sim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sim.Status = "active"

	if err := views.DB.Create(&sim).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to activate sim"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sim activated successfully", "sim": sim})
}

// Deactivate is function which used to Deactivate the sim

func (views *NewDB) Deactivate(c *gin.Context) {
	var sim models.SimcardSchema
	phoneNumber := c.Query("phoneno")

	if err := views.DB.Where("phoneno = ?", phoneNumber).First(&sim).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sim card not found"})
		return
	}

	sim.Status = "inactive"

	if err := views.DB.Save(&sim).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate sim"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sim deactivated successfully", "sim": sim})
}

// handler is function which used to get the simdetails

func (views *NewDB) Handler(c *gin.Context) {
	simNumber := c.Param("simnumber") //you have to provide the simnumber

	var sim models.SimcardSchema

	if err := views.DB.Where("sim_number = ?", simNumber).First(&sim).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sim card not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sim_details": sim})
}
