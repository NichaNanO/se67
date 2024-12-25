package entity

import "gorm.io/gorm"

type ContractType struct {
	gorm.Model
    ID             uint   `gorm:"primaryKey" json:"contract_type_id"`
    ContractName   string `json:"contract_name"`
    MonthlyRent    uint   `json:"monthly_rent"`
    DurationMonths uint   `json:"duration_months"`

    // Relationships
    Contracts []Contract `gorm:"foreignKey:ContractTypeID" json:"contracts"`
}
