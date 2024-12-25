package entity

import "gorm.io/gorm"

type ContractStatus struct {
	gorm.Model
	Status    string
	Contracts []Contract `gorm:"foreignKey:StatusID"` // เพิ่ม foreignKey เพื่อเชื่อมโยงกับ Contracts
}
