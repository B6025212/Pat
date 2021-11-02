package controller

import (
	"github.com/b6025212/sa-64-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /patient
func CreatePatient(c *gin.Context) {
	var patient entity.Patient
	var nurse entity.Nurse
	var medicine entity.Medicine
	var disease entity.Disease

	//bind เข้าตัวแปร patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา disease ด้วยid
	if tx := entity.DB().Where("id = ?", patient.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// ค้นหา nurse ด้วยid
	if tx := entity.DB().Where("id = ?", patient.NurseID).First(&nurse); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}

	// ค้นหา medicine ด้วยid
	if tx := entity.DB().Where("id = ?", patient.MedicineID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video not found"})
		return
	}
	// 12: สร้าง patient
	pa := entity.Patient{
		Identification: patient.Identification,
		Patient_name:   patient.Patient_name,
		Medicine:       medicine,     // โยงความสัมพันธ์กับ Entity medicine
		Disease:        disease,      // โยงความสัมพันธ์กับ Entity disease
		Nurse:          nurse,        // โยงความสัมพันธ์กับ Entity nurse
		Date:           patient.Date, // ตั้งค่าฟิลด์ watchedTime
	}

	// 13: บันทึก
	if err := entity.DB().Create(&pa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pa})
}

// GET /patient/:id
func GetPatient(c *gin.Context) {
	var patient entity.Patient
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", id).Scan(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// GET /patients
func ListPatients(c *gin.Context) {
	var patients []entity.Patient
	if err := entity.DB().Raw("SELECT * FROM patients").Scan(&patients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// DELETE /patient/:id
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patients WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /patient
func UpdatePatient(c *gin.Context) {
	var patient entity.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", patient.ID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}
	if err := entity.DB().Save(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patient})
}
