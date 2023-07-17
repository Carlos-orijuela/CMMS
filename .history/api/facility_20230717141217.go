package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	fa := models.Facility{}
	name := r.FormValue("name")
	description := r.FormValue("dscrptn")
	parentFacility := r.FormValue("pfa")
	code := r.FormValue("code")
	category := r.FormValue("cat")
	address := r.FormValue("addr")
	googlemaps := r.FormValue("gmaps")
	city := r.FormValue("city")
	country := r.FormValue("cntry")

	fa.Name = name
	fa.Description = description
	fa.ParentFacility = parentFacility
	fa.Code = code
	fa.Category = category
	fa.Address = address
	fa.GoogleMaps = googlemaps
	fa.City = city
	fa.Country = country

	db.Save(&fa)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	item := []models.Facility{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	fa := models.Facility{}

	name := r.FormValue("name")
	description := r.FormValue("dscrptn")
	parentFacility := r.FormValue("pfa")
	code := r.FormValue("code")
	category := r.FormValue("cat")
	address := r.FormValue("addr")
	googlemaps := r.FormValue("gmaps")
	city := r.FormValue("city")
	country := r.FormValue("cntry")

	id, _ := strconv.Atoi(r.FormValue("id"))
	db.Where("id", id).Find(&fa)

	fa.Name = name
	fa.Description = description
	fa.ParentFacility = parentFacility
	fa.Code = code
	fa.Category = category
	fa.Address = address
	fa.GoogleMaps = googlemaps
	fa.City = city
	fa.Country = country

	db.Save(&fa)

}

func DeleteFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Facility{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
