package main

import (
	"log"
	"net/http"
	"techno/handlers"
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

const portNumber = ":8080"

func main() {
	http.HandleFunc("/admin/redirects", handlers.GetAdminRedirects)
	http.HandleFunc("/redirects", handlers.UserRedirect)

	log.Fatal(http.ListenAndServe(portNumber, nil))
}
