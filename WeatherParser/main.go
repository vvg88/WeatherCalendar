package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weatherdata/", yaWeatherDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func yaWeatherDataHandler(w http.ResponseWriter, r *http.Request) {
	page, err := readHTTPPage()
	if err != nil {
		log.Fatal(err)
	}
	wd := parsePageAsync(page)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wd)
	wdAsJSON, err := wd.toJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(wdAsJSON))

	err = wd.saveToFireStore()
	if err != nil {
		log.Fatal(err)
	}
}
