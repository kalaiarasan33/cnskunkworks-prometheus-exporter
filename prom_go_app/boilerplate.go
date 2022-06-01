package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	// Start the application
	startMyApp()
}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/birthday/{name}", func(rw http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			name := vars["name"]
			greetings := fmt.Sprintf("Happy Birthday %s :)", name)
			rw.Write([]byte(greetings))
	}).Methods("GET")

	log.Println("Starting the application server...")
	http.ListenAndServe(":8000", router)
}