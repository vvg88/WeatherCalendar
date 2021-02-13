package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

var fsClient *firestore.Client
var ctx context.Context

func main() {
	r := mux.NewRouter()

	ctx = context.Background()
	var err error
	fsClient, err = createClient(ctx)
	defer fsClient.Close()
	if err != nil {
		log.Fatalf("Unable to create firestore client!\nError: %v", err)
	}

	r.HandleFunc("/dbservice/weatherdata", postWeatherData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func postWeatherData(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading POST request body: %v", err)
	}
	log.Println(string(reqBody))

	var wd weatherData
	err = json.Unmarshal(reqBody, &wd)
	if err != nil {
		log.Printf("Error parsing POST json: %v", err)
	}
	log.Println(wd)

	fswd := wd.toFirestoreData()
	log.Println(fswd)
	err = fswd.save()
	if err != nil {
		log.Printf("Error saving weather data to FireStore: %v", err)
	}
	log.Println("Value saved!")
}
