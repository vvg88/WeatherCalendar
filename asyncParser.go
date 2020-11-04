package main

import (
	"io/ioutil"
	"log"
	"time"
)

func getTemperatureAsync(page []byte, c chan int) {
	temp, _ := getTemperature(page)
	c <- temp
}

func getWindSpeedAsync(page []byte, c chan float32) {
	ws, _ := getWindSpeed(page)
	c <- ws
}

func getWindDirectionAsync(page []byte, c chan string) {
	wd, _ := getWindDirection(page)
	c <- wd
}

func getWeatherConditionAsync(page []byte, c chan string) {
	wc, _ := getWeatherCondition(page)
	c <- wc
}

func getHumidityAsync(page []byte, c chan uint8) {
	h, _ := getHumidity(page)
	c <- h
}

func getAirPressureAsync(page []byte, c chan uint16) {
	ap, _ := getAirPressure(page)
	c <- ap
}

func parsePageAsync(name string) *WeatherData {
	page, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	tc := make(chan int)
	wsc := make(chan float32)
	wdc := make(chan string)
	wcc := make(chan string)
	hc := make(chan uint8)
	apc := make(chan uint16)
	go getTemperatureAsync(page, tc)
	go getWindSpeedAsync(page, wsc)
	go getWindDirectionAsync(page, wdc)
	go getWeatherConditionAsync(page, wcc)
	go getHumidityAsync(page, hc)
	go getAirPressureAsync(page, apc)

	return &WeatherData{
		Temperature:      <-tc,
		WindSpeed:        <-wsc,
		WindDirection:    <-wdc,
		WeatherCondition: <-wcc,
		Humidity:         <-hc,
		AirPressure:      <-apc,
		TimeStamp:        time.Now(),
	}
}
