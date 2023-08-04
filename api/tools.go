package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/CMMS/models"
)

func CreateTools(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	fa := models.Tools{}
	fachild := models.ToolsList{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	id := r.FormValue("id")
	dateu := r.FormValue("dateu")
	datep := r.FormValue("datep")
	model := r.FormValue("model")
	serialnumber := r.FormValue("serialnumber")

	if process == "new" {
		
		fachild.Name = name
		fachild.Description = description
		fachild.DatePurchase = datep
		fachild.DateUsed = dateu
		fachild.Model = model
		fachild.SerialNumber = serialnumber
		fachild.EquipmentListID = id

		db.Save(&fachild)
	} else {
		fa.Name = name
		fa.Description = description
		fa.DatePurchase = datep
		fa.DateUsed = dateu
		fa.Model = model
		fa.SerialNumber = serialnumber
		fa.EquipmentID = id
		db.Save(&fa)

	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetTools(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	item := []models.Tools{}
	db.Preload("Equipment").Find(&item)

	itemlist := []models.ToolsList{}
	db.Preload("EquipmentList").Find(&itemlist)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"tools" : itemlist,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func EditTools(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	fa := models.Tools{}
	fachild := models.ToolsList{}
	process := r.FormValue("process")
	name := r.FormValue("name")
	description := r.FormValue("description")
	dateu := r.FormValue("dateu")
	datep := r.FormValue("datep")
	model := r.FormValue("model")
	serialnumber := r.FormValue("serialnumber")
	pid, _ := strconv.Atoi(r.FormValue("pid"))

	if process == "new" {
		db.Where("id", pid).Find(&fachild)
		fachild.Name = name
		fachild.Description = description
		fachild.DatePurchase = datep
		fachild.DateUsed = dateu
		fachild.Model = model
		fachild.SerialNumber = serialnumber
		db.Save(&fachild)
	} else {
		db.Where("id", pid).Find(&fa)
		fa.Name = name
		fa.Description = description
		fa.DatePurchase = datep
		fa.DateUsed = dateu
		fa.Model = model
		fa.SerialNumber = serialnumber
		db.Save(&fa)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func DeleteTools(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	process := r.FormValue("process")

	if (process == "new"){
		item := models.ToolsList{}
		db.Where("id", id).Statement.Delete(&item)
	}else{
		itemlist := models.Tools{}
		db.Where("id", id).Statement.Delete(&itemlist)
	}
	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
