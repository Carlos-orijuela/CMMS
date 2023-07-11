package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreatePosition(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	depart := models.Department{}
	name := r.FormValue("name")

	depart.Name = name
	db.Save(&depart)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetPosition(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	item := []models.Department{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditPosition(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Department{}
	name := r.FormValue("name")
	id, _ := strconv.Atoi(r.FormValue("id"))

	db.Where("id", id).Find(&product)
	product.Name = name
	db.Save(&product)

}

func DeletePositino(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Position{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
