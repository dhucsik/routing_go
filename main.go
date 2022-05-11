package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhucsik/technodom_case_go/handlers"
	"github.com/gorilla/mux"
)

type Cache interface {
	Add(key, value string)
	Get(key string) (value string, ok bool)
	Len() int
}

type Link struct {
	ActiveLink  string `json:"active_link"`
	HistoryLink string `json:"history_link"`
}

const (
	portNumber = ":8080"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/admin/redirects", handlers.GetAdminRedirects).Methods("GET")
	router.HandleFunc("/admin/redirects", handlers.PostAdminRedirects).Methods("POST")
	router.HandleFunc("/admin/redirects/{id:[0-9]+}", handlers.GetAdminRedirectsId).Methods("GET")
	router.HandleFunc("/admin/redirects/{id:[0-9]+}", handlers.PatchAdminRedirectsId).Methods("PATCH")
	router.HandleFunc("/admin/redirects/{id:[0-9]+}", handlers.DeleteAdminRedirectId).Methods("DELETE")

	router.HandleFunc("/redirects", handlers.UserRedirect).Methods("GET")

	fmt.Println("Starting server at", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, router))
}
