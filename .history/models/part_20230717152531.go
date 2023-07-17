package models

import "gorm.io/gorm"

type Parts struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	Image            string
	Commisioneddate  string
	DCommisioneddate string
	ParentFacility   string
	QRCode           string
	Manufacturer     string
	Model            string
	SerialNumber     string
	RunningHr        string
	CriticalY_N      bool // YES OR NO ??? bool or string???

	Deleted gorm.DeletedAt
}

// ?EQUIPMENT RELATIONSHIP
// *FACILITY TABLE
// *LOCATION TABLE
// *EQUIPMENTS TABLE
// *SUPPLIERS TABLE
