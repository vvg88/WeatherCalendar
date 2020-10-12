package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

const tmpExpr = `Текущая температура<\/span><span class="temp__value">([+-]\d+)<\/span>`
const windSpeedExpr = `wind-speed">(\d{1,2}(,\d)?)<\/span>`
const wethCondExpr = `link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":\d+}}'>([а-яА-Я\s]+)`
const humExpr = `icon_humidity-white term__fact-icon" aria-hidden="true"><\/i>(\d{1,2})`
const pressExpr = `icon_pressure-white term__fact-icon" aria-hidden="true"><\/i>(?P<press>\d{3})`

var tempRegx = regexp.MustCompile(tmpExpr)
var windSpeedRegx = regexp.MustCompile(windSpeedExpr)
var conditionRegx = regexp.MustCompile(wethCondExpr)
var humidRegx = regexp.MustCompile(humExpr)
var pressRegx = regexp.MustCompile(pressExpr)

func parsePage(name string) {
	page, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	submatches := tempRegx.FindSubmatch(page)
	if submatches == nil {
		fmt.Println("No temperature matches found!")
	}
	fmt.Printf("A current temperature is %s\n", string(submatches[1]))

	submatches = windSpeedRegx.FindSubmatch(page)
	if submatches == nil {
		fmt.Println("No wind speed matches found!")
	}
	fmt.Printf("A wind speed is %s\n", string(submatches[1]))

	submatches = conditionRegx.FindSubmatch(page)
	if submatches == nil {
		fmt.Println("No weather condition found!")
	} else {
		fmt.Printf("Weather condition: %s\n", string(submatches[1]))
	}

	submatches = humidRegx.FindSubmatch(page)
	if submatches == nil {
		fmt.Println("No humidity found!")
	}
	fmt.Printf("Humidity: %s", string(submatches[1]))
	fmt.Println("%")

	submatches = pressRegx.FindSubmatch(page)
	if submatches == nil {
		fmt.Println("No air pressure found!")
	}
	fmt.Printf("Air pressure: %s", string(submatches[1]))
}

func readPage(name string) (string, error) {
	page, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(page), err
}
