package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	fa := models.Facility{}
	fachild := models.ChildFacility{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	id := r.FormValue("id")
	code := r.FormValue("123")
	category := r.FormValue("category")
	location := r.FormValue("location")
	googlemaps := r.FormValue("gmaps")
	runninghr := r.FormValue("running")

	if process == "new" {
		fa.Name = name
		fa.Description = description
		fa.Code = code
		fa.Category = category
		fa.GoogleMaps = googlemaps
		fa.RunningHr = runninghr
		fa.LocationListID = location

		db.Save(&fa)
	} else {
		fachild.Name = name
		fachild.Description = description
		fachild.FacilityID = id
		fachild.Code = code
		db.Save(&fachild)

	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	item := []models.Facility{}
	db.Preload("LocationList").Find(&item)

	itemlist := []models.ChildFacility{}
	db.Preload("Facility").Find(&itemlist)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"facility" : itemlist,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	fa := models.Facility{}
	fachild := models.ChildFacility{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	id := r.FormValue("id")
	code := r.FormValue("123")
	category := r.FormValue("category")
	location := r.FormValue("location")
	googlemaps := r.FormValue("gmaps")
	runninghr := r.FormValue("running")
	pid, _ := strconv.Atoi(r.FormValue("pid"))

	if process == "new" {
		db.Where("id", pid).Find(&fa)
		fa.Name = name
		fa.Description = description
		fa.Code = code
		fa.Category = category
		fa.GoogleMaps = googlemaps
		fa.RunningHr = runninghr
		fa.LocationListID = location

		db.Save(&fa)
	} else {
		db.Where("id", pid).Find(&fachild)
		fachild.Name = name
		fachild.Description = description
		fachild.FacilityID = id
		fachild.Code = code
		db.Save(&fachild)

	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func DeleteFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	process := r.FormValue("process")

	if (process == "new"){
		item := models.Facility{}
		db.Where("id", id).Statement.Delete(&item)
		itemlist := models.ChildFacility{}
		db.Where("facility_id", id).Statement.Delete(&itemlist)
	}else{
		itemlist := models.ChildFacility{}
		db.Where("id", id).Statement.Delete(&itemlist)
	}
	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
