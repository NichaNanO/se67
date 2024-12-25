package entity

import "time"

import "gorm.io/gorm"

type ContractDocument struct {
	gorm.Model
    ID         uint      `gorm:"primaryKey" json:"document_id"`
    FileName   string    `json:"file_name"`
    UploadDate time.Time `json:"upload_date"`
    FileURL    string    `json:"file_url"`
    Signature  string    `json:"signature"` // เพิ่มฟิลด์สำหรับลายเซ็น (Base64)

    // Foreign Key
    ContractID uint `json:"contract_id"`

    // Relationships
    Contract Contract `gorm:"foreignKey:ContractID" json:"contract"`
}

