package models

import "gorm.io/gorm"

type Tools struct {
	ID            uint `gorm:"primaryKey"`
	Equipment	  Equipment
	EquipmentID	  string
	Name          string
	Description   string
	Model		  string
	SerialNumber  string
	DatePurchase  string
	DateUsed      string
	Deleted       gorm.DeletedAt
}


type ToolsList struct {
	ID            uint `gorm:"primaryKey"`
	EquipmentList	  EquipmentList
	EquipmentListID	  string
	Name          	  string
	Description       string
	Model		      string
	SerialNumber      string
	DatePurchase      string
	DateUsed          string
	Deleted           gorm.DeletedAt
}

// ?EQUIPMENT RELATIONSHIP
// *FACILITY TABLE
// *LOCATION TABLE
// *EQUIPMENTS TABLE
// *SUPPLIERS TABLE
