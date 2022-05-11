package handlers

import (
	"fmt"
	"net/http"
)

func GetAdminRedirects(w http.ResponseWriter, r *http.Request) {
	// get list of sthing from db
	fmt.Fprint(w, "Get method received.")
}

func PostAdminRedirects(w http.ResponseWriter, r *http.Request) {
	// create new sthing in db
	fmt.Fprint(w, "Post method received.")
}

func GetAdminRedirectsId(w http.ResponseWriter, r *http.Request) {
	//get specified object
	fmt.Fprint(w, "Get method received")
}

func PatchAdminRedirectsId(w http.ResponseWriter, r *http.Request) {
	//
	fmt.Fprint(w, "Patch method received")
}

func DeleteAdminRedirectId(w http.ResponseWriter, r *http.Request) {
	// delete sthing from db
	fmt.Fprint(w, "Delete method received")
}

func UserRedirect(w http.ResponseWriter, r *http.Request) {
	// do something
	fmt.Fprint(w, "Get method received.")
}
