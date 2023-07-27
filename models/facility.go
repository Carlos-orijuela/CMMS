package models

import "gorm.io/gorm"

type Facility struct {
	ID             uint `gorm:"primaryKey"`
	LocationList	LocationList
	LocationListID	string
	Name           string
	Description    string
	Code           string // ?AUTO GENERATE
	Category       string // NOTE: DDL >> Storage, Plants, Building
	RunningHouse   string // ?YES OR NO
	RunningHr      string
	GoogleMaps     string
	Deleted        gorm.DeletedAt
}

type ChildFacility struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Description    string
	Facility   Facility
	FacilityID string
	Code       string
	Deleted    gorm.DeletedAt
}

// ?FACILITY RELATIONSHIP
// *LOCATION TABLE
// *EQUIPMENT TABLE
// *STOOLS TABLE
// *PARTS TABLE
