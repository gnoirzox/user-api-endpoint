package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./locations"
	"./users"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", users.PostUser).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", users.GetUser).Methods("GET")
	router.HandleFunc("/user/location", locations.PostLocation).Methods("POST")

	log.Fatal(http.ListenAndServe(":8888", router))
}
