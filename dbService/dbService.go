package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/dbservice/weatherdata", postWeatherData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func postWeatherData(w http.ResponseWriter, r *http.Request) {

}
