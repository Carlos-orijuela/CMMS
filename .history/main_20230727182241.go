package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dafalo/CMMS/api"
	"github.com/dafalo/CMMS/models"
	"github.com/dafalo/CMMS/views"

	// hash password
	"golang.org/x/crypto/bcrypt"

	// mysql
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":3000"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	CreateDB("CMMS")
	MigrateDB()
	CreateGroup()
	CreateDefaultUser()
	CreatePosition()
	Handlers()

	http.ListenAndServe(Port, nil)
}

func Handlers() {
	fmt.Println("Handlers")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", views.LoginHandler)
	http.HandleFunc("/api/", api.APIHandler)

	http.HandleFunc("/dashboard", views.DashboardHandler)
	http.HandleFunc("/clientdashboard", views.ClientDashboardHandler)

	http.HandleFunc("/users", views.UserHandler)
	http.HandleFunc("/group", views.GroupHandler)
	http.HandleFunc("/position", views.SystemRoleHandler)
	http.HandleFunc("/permission", views.PermissionHandler)

	http.HandleFunc("/facility", views.FacilityHandler)
	http.HandleFunc("/location", views.LocationHandler)
	http.HandleFunc("/equipment", views.LocationHandler)

}

func CreateDB(name string) *sql.DB {
	fmt.Println("Database Created")

	//user:password ======= root:a
	db, err := sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func MigrateDB() {
	fmt.Println("Database Migrated")
	user := models.User{}
	group := models.Group{}
	permission := models.Permission{}
	permlist := models.Permissionlist{}
	grouplist := models.Grouplist{}
	facility := models.Facility{}
	facilitylist := models.ChildFacility{}
	location := models.Location{}
	loclist := models.LocationList{}
	db := GormDB()
	db.AutoMigrate(&user, &group, &permission, &permlist, &grouplist, &facility, &facilitylist, &location, &loclist)
}

func CreateGroup() {

	db := GormDB()

	user := []models.Group{}
	db.Find(&user)

	defaultUser := []models.Group{
		{
			Name:        "Default_Group",
			Description: "Description here.",
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Name == users.Name {
				isExisting = true
				break
			}
		}
		if !isExisting {
			fmt.Println("Create Default Group")
			db.Save(&defaultUser[i])
		}
	}

}

func CreatePosition() {

	db := GormDB()

	user := []models.Position{}
	db.Find(&user)

	defaultUser := []models.Position{
		{
			Name:        "Admin",
			Description: "Admin",
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Name == users.Name {
				isExisting = true
				break
			}
		}
		if !isExisting {
			fmt.Println("Create Default Position")
			db.Save(&defaultUser[i])
		}
	}

}

func CreateDefaultUser() {

	db := GormDB()

	user := []models.User{}
	db.Find(&user)

	defaultUser := []models.User{
		{
			Username:   "admin",
			Email:      "dave@gmail.com",
			Password:   hashPassword("admin"),
			Name:       "Dave Falo",
			PositionID: "1",
		},

		{
			Username:   "user",
			Password:   hashPassword("user"),
			Email:      "jw@gmail.com",
			Name:       "John Wick",
			PositionID: "1",
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Username == users.Username {
				isExisting = true
				break
			}
		}
		if !isExisting {
			fmt.Println("Create Default User")
			db.Save(&defaultUser[i])
		}
	}

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func GormDB() *gorm.DB {
	dsn := "root:a@tcp(127.0.0.1:3306)/CMMS?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}
