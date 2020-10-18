package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const yaURL = "https://yandex.ru/pogoda/?lat=56.979999&lon=40.983071"

func main() {
	readYaPage()
	wd := parsePage("yaPage.txt")
	fmt.Println(wd)
	wd.save()
}

func readYaPage() {
	page, err := readHTTPPage()
	err = savePage("yaPage.txt", page)
	if err != nil {
		log.Fatal(err)
	}
}

func readHTTPPage() ([]byte, error) {
	resp, err := http.Get(yaURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return body, err
}

func savePage(file string, page []byte) error {
	err := ioutil.WriteFile(file, page, 0600)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
