package controller

import (
	"net/http"

	"github.com/b6025212/sa-64-example/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// POST /nurse
func CreateNurse(c *gin.Context) {
	var nurse entity.Nurse
	if err := c.ShouldBindJSON(&nurse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(nurse.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	nurse.Password = string(bytes)

	if err := entity.DB().Create(&nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nurse})
}

// GET /nurse/:id
func GetNurse(c *gin.Context) {
	var nurse entity.Nurse
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM nurses WHERE id = ?", id).Scan(&nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nurse})
}

// GET /nurses
func ListNurses(c *gin.Context) {
	var nurses []entity.Nurse
	if err := entity.DB().Raw("SELECT * FROM nurses").Scan(&nurses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nurses})
}
