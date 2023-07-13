package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	group := models.Group{}
	name := r.FormValue("name")
	description := r.FormValue("description")

	group.Name = name
	group.description = description
	db.Save(&group)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	item := []models.Group{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditGroup(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Group{}
	name := r.FormValue("name")
	id, _ := strconv.Atoi(r.FormValue("id"))

	db.Where("id", id).Find(&product)
	product.Name = name
	db.Save(&product)

}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Group{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
