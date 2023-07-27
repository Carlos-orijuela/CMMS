package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	loc := models.Location{}
	loclist := models.LocationList{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	id := r.FormValue("id")
	code := r.FormValue("123")
	category := r.FormValue("category")
	address := r.FormValue("address")
	googlemaps := r.FormValue("gmaps")
	city := r.FormValue("city")
	country := r.FormValue("country")


	

	if (process == "exist"){
		// Add location under location name
		loclist.LocationID = id
		loclist.Name = name
		loclist.Description = description
		loclist.Code = code

		db.Save(&loclist)

	}else{
		// Add new location name
		loc.Name = name
		loc.Description = description
		loc.Code = code
		loc.Category = category
		loc.Address = address
		loc.GoogleMaps = googlemaps
		loc.City = city
		loc.Country = country
	
		db.Save(&loc)
	}

	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	item := []models.Location{}
	db.Find(&item)

	itemlist := []models.LocationList{}
	db.Preload("Location").Find(&itemlist)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"location" : itemlist,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	loc := models.Location{}
	loclist := models.LocationList{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	id := r.FormValue("id")
	code := r.FormValue("123")
	category := r.FormValue("category")
	address := r.FormValue("address")
	googlemaps := r.FormValue("gmaps")
	city := r.FormValue("city")
	country := r.FormValue("country")
	pid, _ := strconv.Atoi(r.FormValue("pid"))


	

	if (process == "exist"){
		// Edit location under location name
		db.Where("id", pid).Find(&loclist)
		loclist.LocationID = id
		loclist.Name = name
		loclist.Description = description
		loclist.Code = code

		db.Save(&loclist)

	}else{
		// Edit new location
		db.Where("id", pid).Find(&loc)
		loc.Name = name
		loc.Description = description
		loc.Code = code
		loc.Category = category
		loc.Address = address
		loc.GoogleMaps = googlemaps
		loc.City = city
		loc.Country = country
	
		db.Save(&loc)
	}
	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	process := r.FormValue("process")

	if (process == "new"){
		item := models.Location{}
		db.Where("id", id).Statement.Delete(&item)
		itemlist := models.LocationList{}
		db.Where("location_id", id).Statement.Delete(&itemlist)
	}else{
		itemlist := models.LocationList{}
		db.Where("id", id).Statement.Delete(&itemlist)
	}
	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
