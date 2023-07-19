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
	if strings.HasPrefix(r.URL.Path, "group") {
		CreateGroup(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_group") {
		GetGroup(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_group") {
		EditGroup(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_group") {
		DeleteGroup(w, r)
		return
	}
	//=====================================================

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
	//=====================================================

	//=====================================================
	if strings.HasPrefix(r.URL.Path, "permission") {
		CreatePermission(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_permission") {
		GetPermission(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_permission") {
		EditPermission(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_permission") {
		DeletePermission(w, r)
		return
	}
	//=====================================================
	if strings.HasPrefix(r.URL.Path, "location") {
		CreateLocation(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_location") {
		GetLocation(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_location") {
		EditLocation(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_location") {
		DeleteLocation(w, r)
		return
	}
	//=====================================================
	if strings.HasPrefix(r.URL.Path, "facility") {
		CreateFacility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_facility") {
		GetFacility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_facility") {
		EditFacility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_facility") {
		DeleteFacility(w, r)
		return
	}

}
