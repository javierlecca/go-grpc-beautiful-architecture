package entity

import (
	"github.com/jinzhu/gorm"
)

type Society struct {
	ID               uint32  `json:"id" db:"id" gorm:"primaryKey; autoIncrement; column:id"`
	Name             string  `json:"name" db:"name" gorm:"varchar(255); column:name"`
	NameComplement   string  `json:"name_complement" db:"name_complement" gorm:"varchar(255); column:name_complement"`
	Logo             string  `json:"logo" db:"logo" gorm:"varchar(255); column:logo"`
	Abstract         string  `json:"abstract" db:"abstract" gorm:"varchar(255); column:abstract"`
	AbstractPhoto    string  `json:"abstract_photo" db:"abstract_photo" gorm:"varchar(255); column:abstract_photo"`
	Country          string  `json:"country" db:"country" gorm:"varchar(255); column:country"`
	Department       string  `json:"department" db:"department" gorm:"varchar(255); column:department"`
	FamiliesNumber   int32   `json:"families_number" db:"families_number" gorm:"int(); column:families_number"`
	AssociatesNumber int32   `json:"associates_number" db:"associates_number" gorm:"int(); column:associates_number"`
	SalesNumber      int32   `json:"sales_number" db:"sales_number" gorm:"int(); column:sales_number"`
	Longitude        float32 `json:"longitude" db:"longitude" gorm:"decimal(); column:longitude"`
	Latitude         float32 `json:"latitude" db:"latitude" gorm:"decimal(); column:latitude"`
	Web              string  `json:"web" db:"web" gorm:"varchar(255); column:web"`
	Video            string  `json:"video" db:"video" gorm:"varchar(255); column:video"`
}

func (Society) TableName() string {
	return "society"
}

func (s *Society) BeforeCreate(scope *gorm.Scope) {
	// uuid := uuid.NewV4()
	// scope.SetColumn("ID", uuid.String())
	// scope.SetColumn("IsDelete", false)
}

type APISocietyListAll struct {
	ID         uint    `json:"id"`
	Abstract   string  `json:"abstract"`
	Country    string  `json:"country"`
	Department string  `json:"department"`
	Longitude  float32 `json:"longitude"`
	Latitude   float32 `json:"latitude"`
	Web        string  `json:"web"`
	Video      string  `json:"video"`
	Person     string  `json:"person"`
	Photo      string  `json:"photo"`
}

// func (SocietyAPI) TableName() string {
// 	return "society"
// }
