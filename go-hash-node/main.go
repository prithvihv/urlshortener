package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	initDb()
	initRedis()
	router := mux.NewRouter().StrictSlash(true)
	// registing routes
	router.HandleFunc("/", createShortURL).Methods("POST")
	router.HandleFunc("/{tinyuid}", getFullURL).Methods("GET")
	fmt.Println("[HTTP] hashnode running at ")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"localhost:8080", "*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
