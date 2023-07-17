package models

import "gorm.io/gorm"

type Equipment struct {
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
	CriticalY_N      bool

	Deleted gorm.DeletedAt
}
