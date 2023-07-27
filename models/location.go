package models

import "gorm.io/gorm"

type Location struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Description    string
	Code           string
	Category       string // remarks : Store Room, Plants, Building
	Address        string
	GoogleMaps     string
	City           string
	Country        string
	Deleted        gorm.DeletedAt
}


type LocationList struct {
	ID             uint `gorm:"primaryKey"`
	Location		Location
	LocationID		string
	Name           string
	Description    string
	Code			string
	Deleted        gorm.DeletedAt
}

// ?LOCATION RELATIONSHIP
// *FACILITY TABLE
