package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.User{}
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	position := r.FormValue("position")
	department := r.FormValue("department")

	product.Name = name
	product.Username = username
	product.Email = email
	product.Password = hashPassword(password)
	product.PositionID = position

	db.Save(&product)

	// * Convert the JSON into an array of map
	var c []map[string]string
	json.Unmarshal([]byte(department), &c)
	// * Declare the struct for the passing of values
	// * Save Each Permission ID
	for i := range c {
		permlist := models.Grouplist{}

		permlist.UserID = fmt.Sprint(product.ID)
		permlist.GroupID = c[i]["groupID"]
		db.Save(&permlist)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.User{}
	db.Preload("Position").Find(&item)
	group := []models.Grouplist{}
	db.Preload("User").Preload("Group").Find(&group)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"group":  group,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	product := models.User{}
	name := r.FormValue("name")
	username := r.FormValue("username")
	email := r.FormValue("email")
	position := r.FormValue("position")
	group := r.FormValue("group")

	db.Where("id", id).Find(&product)

	product.Name = name
	product.Username = username
	product.Email = email
	product.PositionID = position
	db.Save(&product)

	// * Convert the JSON into an array of map
	var c []map[string]string
	json.Unmarshal([]byte(group), &c)
	// * Declare the struct for the passing of values
	// * Save Each Permission ID
	for i := range c {
		permlist := models.Grouplist{}

		permlist.UserID = fmt.Sprint(product.ID)
		permlist.GroupID = c[i]["groupID"]
		db.Save(&permlist)
		db.Model(&permlist).Updates("name", "hello")

	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.User{}
	db.Where("id", id).Statement.Delete(&item)

	group := models.Grouplist{}
	db.Where("user_id", id).Statement.Delete(&group)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}
