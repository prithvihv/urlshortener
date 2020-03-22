package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDb()
	router := mux.NewRouter().StrictSlash(true)
	// registing routes
	router.HandleFunc("/{url}", createShortURL).Methods("POST")
	router.HandleFunc("/", getFullURL).Methods("GET")
	fmt.Println("[HTTP] hashnode running at ")
	log.Fatal(http.ListenAndServe(":9000", router))
}
