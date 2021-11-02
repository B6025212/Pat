package controller

import (
	"github.com/b6025212/sa-64-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /disease
func CreateDisease(c *gin.Context) {
	var disease entity.Disease
	if err := c.ShouldBindJSON(&disease); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": disease})
}

// GET /disease/:id
func GetDisease(c *gin.Context) {
	var disease entity.Disease
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM diseases WHERE id = ?", id).Scan(&disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": disease})
}

// GET /diseases
func ListDiseases(c *gin.Context) {
	var diseases []entity.Disease
	if err := entity.DB().Raw("SELECT * FROM diseases").Scan(&diseases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diseases})
}

// DELETE /disease/:id
func DeleteDisease(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM diseases WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /disease
func UpdateDisease(c *gin.Context) {
	var disease entity.Disease
	if err := c.ShouldBindJSON(&disease); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", disease.ID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}
	if err := entity.DB().Save(&disease).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": disease})
}
