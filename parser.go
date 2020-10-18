package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func getTemperature(page []byte) (int, error) {
	submatches := tempRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No temperature matches found!")
		return 0, nil
	}
	tempStr := string(submatches[1])
	t, err := strconv.Atoi(tempStr)
	if err != nil {
		log.Printf("Unable to parse temperature value: %s\n", tempStr)
		log.Println(err)
		return 0, err
	}
	log.Printf("A current temperature is %d\n", t)
	return t, nil
}

func getWindSpeed(page []byte) (float32, error) {
	submatches := windSpeedRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No wind speed matches found!")
		return 0, nil
	}
	wsMatch := bytes.Replace(submatches[1], []byte{','}, []byte{'.'}, 1)
	wsStr := string(wsMatch)
	wsStr = strings.ReplaceAll(wsStr, ",", ".")
	ws, err := strconv.ParseFloat(wsStr, 32)
	if err != nil {
		log.Printf("Unable to parse wind speed value: %s\n", wsStr)
		log.Println(err)
		return 0, err
	}
	log.Printf("A current wind speed is %f\n", ws)
	return float32(ws), nil
}

func getWeatherCondition(page []byte) (string, error) {
	submatches := conditionRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No weather condition matches found!")
		return "", nil
	}
	wcStr := string(submatches[1])
	log.Printf("A current weather condition is %s\n", wcStr)
	return wcStr, nil
}

func getHumidity(page []byte) (uint8, error) {
	submatches := humidRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No humidity matches found!")
		return 0, nil
	}
	humStr := string(submatches[1])
	hum, err := strconv.Atoi(humStr)
	if err != nil {
		log.Printf("Unable to parse humidity value: %s\n", humStr)
		log.Println(err)
		return 0, err
	}
	log.Printf("A current humidity is %d\n", hum)
	return uint8(hum), nil
}

func getAirPressure(page []byte) (uint16, error) {
	submatches := pressRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No air pressure matches found!")
		return 0, nil
	}
	apStr := string(submatches[1])
	ap, err := strconv.Atoi(apStr)
	if err != nil {
		log.Printf("Unable to parse air pressure value: %s\n", apStr)
		log.Println(err)
		return 0, err
	}
	log.Printf("A current air pressure is %d\n", ap)
	return uint16(ap), nil
}

func parsePage(name string) *WeatherData {
	page, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	t, _ := getTemperature(page)
	ws, _ := getWindSpeed(page)
	wc, _ := getWeatherCondition(page)
	h, _ := getHumidity(page)
	ap, _ := getAirPressure(page)
	ts := time.Now()
	return &WeatherData{Temperature: t, WindSpeed: ws, WeatherCondition: wc, Humidity: h, AirPressure: ap, TimeStamp: ts}
}
