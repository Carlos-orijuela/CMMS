package api

import (
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
	stat := r.FormValue("position")

	product.Name = name
	product.Username = username
	product.Password = hashPassword(password)
	product.Type = stat
	product.OTP = "0"

	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.User{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	product := models.User{}
	username := r.FormValue("username")
	password := r.FormValue("password")

	db.Where("username", username).Find(&product)
	product.Password = hashPassword(password)
	product.Default = "1"
	db.Save(&product)

}

func OTP(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	product := models.User{}
	otp := r.FormValue("otp")
	user, _ := r.Cookie("id")

	db.Where("id", user.Value).Find(&product)
	product.OTP = otp
	db.Save(&product)

}

func EditUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	product := models.User{}
	name := r.FormValue("name")
	stat := r.FormValue("type")
	username := r.FormValue("username")
	password := r.FormValue("password")
	action := r.FormValue("action")

	db.Where("id", id).Find(&product)

	if action == "edit" {
		product.Name = name
		product.Username = username
		product.Password = hashPassword(password)
	} else {
		product.Name = name
		product.Type = stat
	}

	db.Save(&product)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.User{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}
