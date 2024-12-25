package config

import (
	"fmt"
	"time"

	"example.com/project/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	// AutoMigrate สำหรับ entity ทั้งหมดที่เกี่ยวข้องกับ contract
	db.AutoMigrate(
		&entity.Users{},
		&entity.Genders{},
		&entity.Contract{},
		&entity.ContractStatus{},
		&entity.ContractType{},
		&entity.ContractDocument{},
		)

	// สร้างข้อมูล Gender
	GenderMale := entity.Genders{Gender: "Male"}
	GenderFemale := entity.Genders{Gender: "Female"}
	db.FirstOrCreate(&GenderMale, &entity.Genders{Gender: "Male"})
	db.FirstOrCreate(&GenderFemale, &entity.Genders{Gender: "Female"})

	// การแฮชรหัสผ่าน
	hashedPassword, _ := HashPassword("123456")

	// การแปลงวันที่
	BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")

	// สร้างหรืออัปเดตผู้ใช้งาน
	User := &entity.Users{
		FirstName: "Software",
		LastName:  "Analysis",
		Email:     "sa@gmail.com",
		Age:       80,
		Password:  hashedPassword,
		BirthDay:  BirthDay,
		GenderID:  1,
	}

	// การสร้างหรือค้นหาผู้ใช้งาน
	db.FirstOrCreate(User, &entity.Users{Email: "sa@gmail.com"})

	// สร้างข้อมูล ContractStatus
	ContractStatusActive := entity.ContractStatus{Status: "Active"}
	ContractStatusExpired := entity.ContractStatus{Status: "Expired"}
	db.FirstOrCreate(&ContractStatusActive, &entity.ContractStatus{Status: "Active"})
	db.FirstOrCreate(&ContractStatusExpired, &entity.ContractStatus{Status: "Expired"})

	// สร้างข้อมูล ContractType
	ContractTypeStandard := entity.ContractType{
		ContractName:   "Standard",
		MonthlyRent:    5000,
		DurationMonths: 12,
	}
	ContractTypeVIP := entity.ContractType{
		ContractName:   "VIP",
		MonthlyRent:    8000,
		DurationMonths: 12,
	}
	db.FirstOrCreate(&ContractTypeStandard, &entity.ContractType{ContractName: "Standard"})
	db.FirstOrCreate(&ContractTypeVIP, &entity.ContractType{ContractName: "VIP"})

	// สร้างข้อมูล Contract
	contract := &entity.Contract{
		StartDate:      time.Now(),
		EndDate:        time.Now().AddDate(0, 12, 0), // 12 เดือนจากปัจจุบัน
		SecurityDeposit: 10000,
		Note:           "This is a sample contract.",
		MemberID:       1, // ใช้ MemberID ที่มีอยู่
		EmployeeID:     User.ID,
		RoomID:         1, // สมมติว่า RoomID มีค่า 1
		ContractTypeID: ContractTypeStandard.ID,
		StatusID:       ContractStatusActive.ID,
	}
	db.FirstOrCreate(contract, &entity.Contract{
		MemberID:   1,
		EmployeeID: User.ID,
	})

	// สร้างข้อมูล ContractDocument
	contractDocument := &entity.ContractDocument{
		FileName:   "contract_sample.pdf",
		UploadDate: time.Now(),
		FileURL:    "/uploads/contract_sample.pdf",
		ContractID: contract.ID,
	}
	db.FirstOrCreate(contractDocument, &entity.ContractDocument{
		ContractID: contract.ID,
	})
}
