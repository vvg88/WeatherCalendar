package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const yaURL = "https://yandex.ru/pogoda/?lat=56.979999&lon=40.983071"
const yaPageFile = "yaPage.txt"

func readAndSaveYaPage() {
	page, err := readHTTPPage()
	if err != nil {
		log.Println(err)
		return
	}
	err = savePage(yaPageFile, page)
	if err != nil {
		log.Fatal(err)
	}
}

func readHTTPPage() ([]byte, error) {
	resp, err := http.Get(yaURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func savePage(file string, page []byte) error {
	err := ioutil.WriteFile(file, page, 0600)
	if err != nil {
		return err
	}
	return nil
}

func openPage(file string) ([]byte, error) {
	page, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return page, nil
}
