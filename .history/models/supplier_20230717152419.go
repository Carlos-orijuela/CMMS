package models

import "gorm.io/gorm"

type Supplier struct {
	ID            uint `gorm:"primaryKey"`
	CompanyName   string
	ContactPerson string
	Category      string //? Category: Supplier, Manufacturer or Service Provider
	Website       string
	Mobile        string
	Landline      string
	Address       string
	Remarks       string
	Deleted       gorm.DeletedAt
}

// ?SUPPLIER RELATIONSHIP
// *EQUIPMENT TABLE
// *STOOLS TABLE
// *PARTS TABLE
