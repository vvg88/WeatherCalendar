package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	readYaPage()
}

func readYaPage() {
	page, err := readHttpPage("https://yandex.ru/pogoda/?lat=56.979999&lon=40.983071")
	err = savePage("yaPage.txt", page)
	if err != nil {
		log.Fatal(err)
	}
}

func readHttpPage(url string) ([]byte, error) {
	resp, err := http.Get(url)
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
