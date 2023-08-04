package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	loc := models.Equipment{}
	loclist := models.EquipmentList{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	facility := r.FormValue("facility")
	subfacility := r.FormValue("subfacility")
	id := r.FormValue("id")
	code := r.FormValue("code")
	manufacturer := r.FormValue("manufacturer")
	model := r.FormValue("model")
	serialnumber := r.FormValue("serialnumber")
	runninghr := r.FormValue("runninghr")
	critical := r.FormValue("critical")
	comdate := r.FormValue("comdate")
	decomdate := r.FormValue("decomdate")
	


	

	if (process == "exist"){
		// Add location under location name
		loclist.EquipmentID = id
		loclist.Name = name
		loclist.Description = description
		loclist.Code = code
		loclist.Manufacturer = manufacturer
		loclist.Model = model
		loclist.SerialNumber = serialnumber
		loclist.RunningHr = runninghr
		loclist.Critical = critical
		loclist.Commisioneddate = comdate
		loclist.DCommisioneddate = decomdate

		db.Save(&loclist)

	}else{
		// Add new location name
		loc.Name = name
		loc.Description = description
		loc.FacilityID = facility
		if (subfacility == ""){
			loc.ChildFacilityID = "1"
		}else{
			loc.ChildFacilityID = subfacility
		}
	
		loc.Code = code
		loc.Manufacturer = manufacturer
		loc.Model = model
		loc.SerialNumber = serialnumber
		loc.RunningHr = runninghr
		loc.Critical = critical
		loc.Commisioneddate = comdate
		loc.DCommisioneddate = decomdate
	
		db.Save(&loc)
	}

	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	item := []models.Equipment{}
	db.Preload("Facility").Preload("ChildFacility").Find(&item)

	itemlist := []models.EquipmentList{}
	db.Preload("Equipment").Find(&itemlist)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"equipment" : itemlist,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	loc := models.Equipment{}
	loclist := models.EquipmentList{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	facility := r.FormValue("facility")
	id := r.FormValue("id")
	code := r.FormValue("code")
	manufacturer := r.FormValue("manufacturer")
	model := r.FormValue("model")
	serialnumber := r.FormValue("serialnumber")
	runninghr := r.FormValue("runninghr")
	critical := r.FormValue("critical")
	comdate := r.FormValue("comdate")
	decomdate := r.FormValue("decomdate")
	pid, _ := strconv.Atoi(r.FormValue("pid"))


	

	if (process == "exist"){
		// Edit location under location name
		db.Where("id", pid).Find(&loclist)
		loclist.EquipmentID = id
		loclist.Name = name
		loclist.Description = description
		loclist.Code = code
		loclist.Manufacturer = manufacturer
		loclist.Model = model
		loclist.SerialNumber = serialnumber
		loclist.RunningHr = runninghr
		loclist.Critical = critical
		loclist.Commisioneddate = comdate
		loclist.DCommisioneddate = decomdate

		db.Save(&loclist)

	}else{
		// Edit new location
		db.Where("id", pid).Find(&loc)
		loc.FacilityID = facility
		loc.Name = name
		loc.Description = description
		loc.Code = code
		loc.Manufacturer = manufacturer
		loc.Model = model
		loc.SerialNumber = serialnumber
		loc.RunningHr = runninghr
		loc.Critical = critical
		loc.Commisioneddate = comdate
		loc.DCommisioneddate = decomdate
		db.Save(&loc)
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func DeleteEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	process := r.FormValue("process")

	if (process == "new"){
		item := models.Equipment{}
		db.Where("id", id).Statement.Delete(&item)
		itemlist := models.EquipmentList{}
		db.Where("equipment_id", id).Statement.Delete(&itemlist)
	}else{
		itemlist := models.EquipmentList{}
		db.Where("id", id).Statement.Delete(&itemlist)
	}
	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
