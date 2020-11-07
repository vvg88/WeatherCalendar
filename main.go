package main

import (
	"fmt"
	"log"
)

func main() {
	page, err := readHTTPPage()
	if err != nil {
		log.Fatal(err)
	}
	wd := parsePageAsync(page)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wd)
	wd.save()
}
