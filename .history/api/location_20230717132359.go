package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	loc := models.Location{}
	name := r.FormValue("name")
	description := r.FormValue("dscrptn")
	parentLocation := r.FormValue("ploc")
	code := r.FormValue("code")
	category := r.FormValue("cat")
	address := r.FormValue("addr")
	googlemaps := r.FormValue("gmaps")
	city := r.FormValue("city")
	country := r.FormValue("cntry")

	loc.Name = name
	loc.Description = description
	loc.ParentLocation = parentLocation
	loc.Code = code
	loc.Category = category
	loc.Address = address
	loc.GoogleMaps = googlemaps
	loc.City = city
	loc.Country = country

	db.Save(&loc)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	item := []models.Location{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	loc := models.Location{}

	name := r.FormValue("name")
	description := r.FormValue("dscrptn")
	parentLocation := r.FormValue("ploc")
	code := r.FormValue("code")
	category := r.FormValue("cat")
	address := r.FormValue("addr")
	googlemaps := r.FormValue("gmaps")
	city := r.FormValue("city")
	country := r.FormValue("cntry")

	id, _ := strconv.Atoi(r.FormValue("id"))
	db.Where("id", id).Find(&loc)

	loc.Name = name
	loc.Description = description
	loc.ParentLocation = parentLocation
	loc.Code = code
	loc.Category = category
	loc.Address = address
	loc.GoogleMaps = googlemaps
	loc.City = city
	loc.Country = country

	db.Save(&loc)

}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Location{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
