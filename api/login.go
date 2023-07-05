package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dafalo/CMMS/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := models.User{}
	report := []models.User{}
	// now := time.Now()
	// date := now.Format("2006-01-02")
	// datetime := now.Format("15:04:05")
	// t := fmt.Sprint(time.Now().Nanosecond())

	db.Where("username = ?", username).Find(&user)
	db.Where("username = ?", username).Find(&report)

	println(username, user.ID)

	db.Where("id", user.ID).Find(&user)

	// if (user.Date == date){
	// 	user.LoginTime = date + " " + datetime
	// 	user.OTP = "1";
	// 	db.Save(&user)
	// }else{
	// 	if user.Username == "admin" || user.Username == "user"{

	// 	}else{

	// 		user.LoginTime = date + " " + datetime
	// 		user.Date = date;
	// 		user.OTP = "0"
	// 		user.Code = t[:6]
	// 		operator := models.SMS{}
	// 		operator.Type = "SEMAPHORE"
	// 		operator.SenderName = "Thesis"

	// 		operator.Username = "b6a29200f850ca721fb33b63470473c2"

	// 		operator.Message = t[:6] + " " + "is your One-Time Password(OTP). Valid for 1 day."

	// 		operator.Phone = user.Number
	// 		operator.Sendsms(false)
	// 		db.Save(&user)
	// 	}
	// }

	if CheckPasswordHash(password, user.Password) {
		result := "1"

		newSession := uuid.NewString()

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "session",
			Value: newSession,
		})

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "id",
			Value: fmt.Sprint(user.ID),
		})
		data := map[string]interface{}{
			"status":  "ok",
			"results": result,
			"reports": report,
		}
		ReturnJSON(w, r, data)
	} else {
		result := "0"
		data := map[string]interface{}{
			"status":  "error",
			"results": result,
		}
		ReturnJSON(w, r, data)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func GormDB() *gorm.DB {
	dsn := "root:a@tcp(127.0.0.1:3306)/CMMS?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	w.Write(b)
}
