package models

import "gorm.io/gorm"

type Facility struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Description    string
	Image          string
	Code           string // AUTO GENERATE
	Category       string //NOTE: DDL >> Storage, Plants, Building
	FilesORManuals string
	RunningHouse   string // ??????? YES OR NO
	RunningHr      string
	QRCode         string
	GoogleMaps     string

	Deleted gorm.DeletedAt
}
