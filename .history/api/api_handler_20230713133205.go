package api

import (
	"net/http"
	"strings"
)

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "login") {
		Login(w, r)
		return
	}

	//=====================================================
	if strings.HasPrefix(r.URL.Path, "user") {
		CreateUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_user") {
		GetUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_user") {
		EditUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_user") {
		DeleteUser(w, r)
		return
	}
	//=====================================================

	if strings.HasPrefix(r.URL.Path, "department") {
		CreateDepartment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_department") {
		GetDepartment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_department") {
		EditDepartment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_department") {
		DeleteDepartment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "position") {
		CreatePosition(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_position") {
		GetPosition(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_position") {
		EditPosition(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_position") {
		DeletePositino(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "password") {
		ChangePassword(w, r)
		return
	}

}
