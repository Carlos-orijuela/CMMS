package models

import "gorm.io/gorm"

type Location struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Description    string
	Image          string
	ParentLocation string
	Code           string
	Category       string // remarks : Store Room, Plants, Building
	Address        string
	GoogleMaps     string
	City           string
	Country        string
	Deleted        gorm.DeletedAt
}

// ?LOCATION RELATIONSHIP
// >FACILITY TABLE
