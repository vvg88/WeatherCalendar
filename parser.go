package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

var tempRegx = regexp.MustCompile(`Текущая температура<\/span><span class="temp__value">([+-]\d+)<\/span>`)
var windSpeedRegx = regexp.MustCompile(`wind-speed">(\d{1,2}(,\d)?)<\/span>`)

func parsePage(name string) {
	page, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	submatches := tempRegx.FindSubmatch(page)
	if submatches == nil {
		log.Fatal("No temperature matches found!")
	}
	fmt.Printf("A current temperature is %s\n", string(submatches[1]))

	submatches = windSpeedRegx.FindSubmatch(page)
	if submatches == nil {
		log.Fatal("No wind speed matches found!")
	}
	fmt.Printf("A wind speed is %s", string(submatches[1]))
}

func readPage(name string) (string, error) {
	page, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(page), err
}
