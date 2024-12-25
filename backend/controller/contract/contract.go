package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/project/config"

	"example.com/project/entity"
)

func GetContracts(c *gin.Context) {
	var contracts []entity.Contract

	db := config.DB()
	results := db.Preload("Member").Preload("Employee").Preload("Room").Preload("ContractType").Find(&contracts)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

func GetContractByID(c *gin.Context) {
	ID := c.Param("id")
	var contract entity.Contract

	db := config.DB()
	results := db.Preload("Member").Preload("Employee").Preload("Room").Preload("ContractType").First(&contract, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, contract)
}

func CreateContract(c *gin.Context) {
	var contract entity.Contract

	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// Validate foreign keys
	var member entity.Member
	db.First(&member, contract.MemberID)
	if member.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบบัญชีสมาชิก"})
		return
	}

	var employee entity.Employee
	db.First(&employee, contract.EmployeeID)
	if employee.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบพนักงาน"})
		return
	}

	// var room entity.Room
	// db.First(&room, contract.RoomID)
	// if room.ID == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบห้องพัก"})
	// 	return
	// }

	var contractType entity.ContractType
	db.First(&contractType, contract.ContractTypeID)
	if contractType.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบประเภทสัญญา"})
		return
	}

	// Create contract
	if err := db.Create(&contract).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "สร้างสัญญาสำเร็จ"})
}

func UpdateContract(c *gin.Context) {
	var contract entity.Contract
	contractID := c.Param("id")

	db := config.DB()
	result := db.First(&contract, contractID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบสัญญา"})
		return
	}

	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง"})
		return
	}

	result = db.Save(&contract)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "แก้ไขข้อมูลไม่สำเร็จ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "แก้ไขข้อมูลสำเร็จ"})
}

func DeleteContract(c *gin.Context) {
	ID := c.Param("id")

	db := config.DB()
	var contract entity.Contract

	if err := db.First(&contract, ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบสัญญา"})
		return
	}

	if err := db.Delete(&contract).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ลบข้อมูลไม่สำเร็จ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ลบข้อมูลสำเร็จ"})
}
