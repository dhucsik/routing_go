package handlers

import (
	"fmt"
	"net/http"
)

func GetAdminRedirects(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// get list of sthing from db
		fmt.Fprint(w, "Get method received.")
	case http.MethodPost:
		// create new sthing in db
		fmt.Fprint(w, "Post method received.")
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}

func UserRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// do something
		fmt.Fprint(w, "Get method received.")
	} else {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}
}
