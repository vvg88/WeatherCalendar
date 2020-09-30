package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	getFromYandex()
}

func getFromYandex() {
	resp, err := http.Get("https://yandex.ru/pogoda/?lat=56.979999&lon=40.983071")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("yaPage.txt", body, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
