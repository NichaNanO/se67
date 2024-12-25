package entity

import "time"

type Contract struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    StartDate       time.Time `json:"start_date"`
    EndDate         time.Time `json:"end_date"`
    SecurityDeposit uint      `json:"security_deposit"`
    Note            string    `json:"note"`

    // Foreign Keys
    MemberID       uint `json:"member_id"`
    EmployeeID     uint `json:"employee_id"`
    RoomID         uint `json:"room_id"`
    ContractTypeID uint `json:"contract_type_id"`
    StatusID       uint `json:"status_id"`

    // Relationships
    Member       Member       `gorm:"foreignKey:MemberID" json:"member"`
    Employee     Employee     `gorm:"foreignKey:EmployeeID" json:"employee"`
    //Room         Room         `gorm:"foreignKey:RoomID" json:"room"`
    ContractType ContractType `gorm:"foreignKey:ContractTypeID" json:"contract_type"`
    Status       ContractStatus `gorm:"foreignKey:StatusID" json:"status"`

}
