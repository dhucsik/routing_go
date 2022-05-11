package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhucsik/technodom_case_go/setupdb"
	"github.com/gorilla/mux"
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

	decoder := json.NewDecoder(r.Body)

	var post Link
	err := decoder.Decode(&post)

	if err != nil {
		panic(err)
	}

	activeLink := post.ActiveLink
	historyLink := post.HistoryLink

	fmt.Println(activeLink)
	fmt.Println(historyLink)

	var response = JsonResponse{}

	db := setupdb.SetupDB()

	var lastInsertId int
	err = db.QueryRow("INSERT INTO links_table(active_link, history_link) VALUES($1, $2) returning id;", activeLink, historyLink).Scan(&lastInsertId)

	if err != nil {
		panic(err)
	}

	response = JsonResponse{Type: "success", Message: "The record has been inserted successfully!"}

	json.NewEncoder(w).Encode(response)
}

func GetAdminRedirectsId(w http.ResponseWriter, r *http.Request) {
	//get specified object
	fmt.Fprint(w, "Get method received")

	params := mux.Vars(r)

	idd := params["id"]

	db := setupdb.SetupDB()

	row, err := db.Query("SELECT * FROM links_table WHERE id = $1;", idd)

	if err != nil {
		panic(err)
	}

	var record []Link
	for row.Next() {
		var id int
		var activeLink string
		var historyLink string

		err = row.Scan(&id, &activeLink, &historyLink)

		if err != nil {
			panic(err)
		}
		record = append(record, Link{
			Id:          id,
			ActiveLink:  activeLink,
			HistoryLink: historyLink,
		})
	}

	response := JsonResponse{Type: "success", Data: record}

	json.NewEncoder(w).Encode(response)
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
