package controllers

import (
	"net/http"
	"hyperchain/api/models"
	"hyperchain/api/services"

	"github.com/gin-gonic/gin"
)

func AddRecord(c *gin.Context) {
	var input models.MedicalRecord
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := services.AddRecord(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Chaincode invoke failed", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Record added"})
}

func GetRecord(c *gin.Context) {
	id := c.Param("id")

	record, err := services.QueryRecord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}
