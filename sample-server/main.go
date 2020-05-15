package main

import (
	"api/action"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	// simple json result
	router.HandleFunc("/tasks", action.GetAllTask).Methods("GET")
	router.HandleFunc("/task/{id}", action.GetTaskOne).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
