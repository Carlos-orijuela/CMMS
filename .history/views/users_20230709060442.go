package views

import (
	"net/http"
	"text/template"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/users.html"))
	data := map[string]interface{}{}
	data["Title"] = "Register | POS_SYSTEM"
	tmpl.Execute(w, data)
}
