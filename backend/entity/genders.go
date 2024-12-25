package entity

import (
	"gorm.io/gorm"
)

type Genders struct {
	gorm.Model
	Name string
	Gender string `json:"gender"`

	Employees []Employee `gorm:"foreignKey:gender_id"`
	Member []Member `gorm:"foreignKey:gender_id"`
}