package models

import "gorm.io/gorm"

type Supplier struct {
	ID            uint `gorm:"primaryKey"`
	CompanyName   string
	ContactPerson string
	Category      string
	Website       string
	Mobile        string
	Landline      string
	Address       string
	Remarks       string
	CriticalY_N   bool // YES OR NO ??? bool or string???

	Deleted gorm.DeletedAt
}

// ?SUPPLIER RELATIONSHIP
// *EQUIPMENT TABLE
// *STOOLS TABLE
// *PARTS TABLE
