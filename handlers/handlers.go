package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhucsik/technodom_case_go/setupdb"
)

type Link struct {
	Id          int    `json:"id"`
	ActiveLink  string `json:"active_link"`
	HistoryLink string `json:"history_link"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Link `json:"data"`
	Message string `json:"message"`
}

func GetAdminRedirects(w http.ResponseWriter, r *http.Request) {
	// get list of sthing from db
	fmt.Fprint(w, "Get method received.")
	db := setupdb.SetupDB()

	rows, err := db.Query("SELECT * FROM links_table;")

	if err != nil {
		panic(err)
	}

	var records []Link

	for rows.Next() {
		var id int
		var activeLink string
		var historyLink string

		err = rows.Scan(&id, &activeLink, &historyLink)

		if err != nil {
			panic(err)
		}

		records = append(records, Link{
			Id:          id,
			ActiveLink:  activeLink,
			HistoryLink: historyLink,
		})
	}

	response := JsonResponse{Type: "success", Data: records}

	json.NewEncoder(w).Encode(response)
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
