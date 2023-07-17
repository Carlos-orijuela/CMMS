package models

import "gorm.io/gorm"

type Facility struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Description    string
	Image          string
	ParentFacility string
	Code           string // ?AUTO GENERATE
	Category       string // NOTE: DDL >> Storage, Plants, Building
	FilesORManuals string
	RunningHouse   string // ?YES OR NO
	RunningHr      string
	QRCode         string // ?AUTO GENERATE
	GoogleMaps     string
	Deleted        gorm.DeletedAt
}

type ChildFacility struct {
	ID           uint `gorm:"primaryKey"`
	Position     Position
	PositionID   string
	Permission   Permission
	PermissionID string
	Deleted      gorm.DeletedAt
}

// ?FACILITY RELATIONSHIP
// *LOCATION TABLE
// *EQUIPMENT TABLE
// *STOOLS TABLE
// *PARTS TABLE
