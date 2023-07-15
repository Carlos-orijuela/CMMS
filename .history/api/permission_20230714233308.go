package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreatePermission(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	group := models.Permission{}
	description := r.FormValue("description")

	group.Description = description
	db.Save(&group)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetPermission(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	item := []models.Permission{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditPermission(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Permission{}

	id, _ := strconv.Atoi(r.FormValue("id"))
	des := r.FormValue("description")

	db.Where("id", id).Find(&product)
	product.Description = des
	db.Save(&product)

}

func DeletePermission(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Permission{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
