package models

import "gorm.io/gorm"

type Parts struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Description   string
	PartUniNumber string
	SerialNumber  string
	DatePurchase  string
	DateUsed      string
	QtyInhand     string
	QtyMin        string
	Manufacturer  string
	FilesManuals  string
	Image         string
	QRCode        string
	RunningHr     string
	CriticalY_N   bool // YES OR NO ??? bool or string???

	Deleted gorm.DeletedAt
}

// ?EQUIPMENT RELATIONSHIP
// *FACILITY TABLE
// *LOCATION TABLE
// *EQUIPMENTS TABLE
// *SUPPLIERS TABLE
