package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateSystemRole(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	depart := models.SystemRole{}
	name := r.FormValue("name")

	depart.Name = name
	db.Save(&depart)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetSystemRole(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	item := []models.SystemRole{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditSystemrole(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.SystemRole{}
	name := r.FormValue("name")
	id, _ := strconv.Atoi(r.FormValue("id"))

	db.Where("id", id).Find(&product)
	product.Name = name
	db.Save(&product)

}

func DeleteSystemRole(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.SystemRole{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
