package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"example.com/project/entity"
	"example.com/project/config"
)

// GetContractStatuses - ดึงข้อมูลสถานะสัญญาทั้งหมด
func GetContractStatuses(c *gin.Context) {
	var contractStatuses []entity.ContractStatus
	if err := config.DB().Find(&contractStatuses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contractStatuses)
}

// GetContractStatusByID - ดึงข้อมูลสถานะสัญญาโดยใช้ ID
func GetContractStatusByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var contractStatus entity.ContractStatus
	if err := config.DB().First(&contractStatus, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract status not found"})
		return
	}
	c.JSON(http.StatusOK, contractStatus)
}

// CreateContractStatus - เพิ่มข้อมูลสถานะสัญญาใหม่
func CreateContractStatus(c *gin.Context) {
	var contractStatus entity.ContractStatus
	if err := c.ShouldBindJSON(&contractStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB().Create(&contractStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, contractStatus)
}

// UpdateContractStatus - อัปเดตข้อมูลสถานะสัญญา
func UpdateContractStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var contractStatus entity.ContractStatus
	if err := config.DB().First(&contractStatus, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract status not found"})
		return
	}

	if err := c.ShouldBindJSON(&contractStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB().Save(&contractStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contractStatus)
}

// DeleteContractStatus - ลบข้อมูลสถานะสัญญา
func DeleteContractStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := config.DB().Delete(&entity.ContractStatus{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contract status deleted successfully"})
}
