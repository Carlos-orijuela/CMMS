package models

import "gorm.io/gorm"

type Equipment struct {
	ID               uint `gorm:"primaryKey"`
	Facility		 Facility
	FacilityID		 string
	ChildFacility	 ChildFacility
	ChildFacilityID	 string
 	Name             string
	Description      string
	Commisioneddate  string
	DCommisioneddate string
	Code           	 string
	Manufacturer     string
	Model            string
	SerialNumber     string
	RunningHr        string
	Critical	     string 
	Deleted 		 gorm.DeletedAt
}

type EquipmentList struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	Equipment		 Equipment
	EquipmentID		 string
	Description      string
	Commisioneddate  string
	DCommisioneddate string
	Code          	 string
	Manufacturer     string
	Model            string
	SerialNumber     string
	RunningHr        string
	Critical	     string 
	Deleted 		gorm.DeletedAt
}




// ?EQUIPMENT RELATIONSHIP
// *FACILITY TABLE
// *LOCATION TABLE
// *STOOLS TABLE
// *SUPPLIERS TABLE
// *PARTS TABLE
