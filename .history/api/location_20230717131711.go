package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	pos := models.Location{}
	name := r.FormValue("name")
	description := r.FormValue("dscrptn")
	parentLocation := r.FormValue("ploc")
	code := r.FormValue("code")
	category := r.FormValue("cat")
	address := r.FormValue("addr")
	googlemaps := r.FormValue("gmaps")
	city := r.FormValue("city")
	country := r.FormValue("cntry")

	pos.Name = role
	pos.Description = desc

	db.Save(&pos)

	// * Convert the JSON into an array of map
	var c []map[string]string
	json.Unmarshal([]byte(result), &c)
	// * Declare the struct for the passing of values
	// * Save Each Permission ID
	for i := range c {
		permlist := models.Permissionlist{}

		permlist.PositionID = fmt.Sprint(pos.ID)
		permlist.PermissionID = c[i]["permID"]
		db.Save(&permlist)
	}

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

	product := models.Location{}
	name := r.FormValue("name")
	id, _ := strconv.Atoi(r.FormValue("id"))

	db.Where("id", id).Find(&product)
	product.Name = name
	db.Save(&product)

}

func DeleteLocation(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Location{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
