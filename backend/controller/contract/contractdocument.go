package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"example.com/project/entity"
	"example.com/project/config"
)

// GetContractDocuments - ดึงข้อมูลเอกสารสัญญาทั้งหมด
func GetContractDocuments(c *gin.Context) {
	var contractDocuments []entity.ContractDocument
	if err := config.DB().Find(&contractDocuments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contractDocuments)
}

// GetContractDocumentByID - ดึงข้อมูลเอกสารสัญญาโดยใช้ ID
func GetContractDocumentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var contractDocument entity.ContractDocument
	if err := config.DB().First(&contractDocument, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract document not found"})
		return
	}
	c.JSON(http.StatusOK, contractDocument)
}

// CreateContractDocument - เพิ่มข้อมูลเอกสารสัญญาใหม่
func CreateContractDocument(c *gin.Context) {
	var contractDocument entity.ContractDocument
	if err := c.ShouldBindJSON(&contractDocument); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB().Create(&contractDocument).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, contractDocument)
}

// UpdateContractDocument - อัปเดตข้อมูลเอกสารสัญญา
func UpdateContractDocument(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var contractDocument entity.ContractDocument
	if err := config.DB().First(&contractDocument, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract document not found"})
		return
	}

	if err := c.ShouldBindJSON(&contractDocument); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB().Save(&contractDocument).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contractDocument)
}

// DeleteContractDocument - ลบข้อมูลเอกสารสัญญา
func DeleteContractDocument(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := config.DB().Delete(&entity.ContractDocument{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contract document deleted successfully"})
}
