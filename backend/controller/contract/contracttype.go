package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"example.com/project/entity"
	"example.com/project/config"
)

// GetContractTypes - ดึงข้อมูลประเภทสัญญาทั้งหมด
func GetContractTypes(c *gin.Context) {
	var contractTypes []entity.ContractType
	db := config.DB()
	if err := db.Find(&contractTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contractTypes)
}

// GetContractTypeByID - ดึงข้อมูลประเภทสัญญาโดยใช้ ID
func GetContractTypeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var contractType entity.ContractType
	db := config.DB()
	if err := db.First(&contractType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract type not found"})
		return
	}
	c.JSON(http.StatusOK, contractType)
}

// CreateContractType - เพิ่มข้อมูลประเภทสัญญาใหม่
func CreateContractType(c *gin.Context) {
	var contractType entity.ContractType
	if err := c.ShouldBindJSON(&contractType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()
	if err := db.Create(&contractType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, contractType)
}

// UpdateContractType - อัปเดตข้อมูลประเภทสัญญา
func UpdateContractType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var contractType entity.ContractType
	db := config.DB()
	if err := db.First(&contractType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract type not found"})
		return
	}

	if err := c.ShouldBindJSON(&contractType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&contractType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contractType)
}

// DeleteContractType - ลบข้อมูลประเภทสัญญา
func DeleteContractType(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	db := config.DB()
	if err := db.Delete(&entity.ContractType{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contract type deleted successfully"})
}
