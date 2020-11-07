package main

import (
	"log"
	"time"
)

func getTemperatureAsync(page []byte, c chan int) {
	temp, err := getTemperature(page)
	if err != nil {
		log.Printf("Error on getting temperature value! \n%v", err)
	}
	c <- temp
}

func getWindSpeedAsync(page []byte, c chan float32) {
	ws, err := getWindSpeed(page)
	if err != nil {
		log.Printf("Error on getting wind speed value! \n%v", err)
	}
	c <- ws
}

func getWindDirectionAsync(page []byte, c chan string) {
	wd, err := getWindDirection(page)
	if err != nil {
		log.Printf("Error on getting wind direction! \n%v", err)
	}
	c <- wd
}

func getWeatherConditionAsync(page []byte, c chan string) {
	wc, err := getWeatherCondition(page)
	if err != nil {
		log.Printf("Error on getting weather conditions! \n%v", err)
	}
	c <- wc
}

func getHumidityAsync(page []byte, c chan uint8) {
	h, err := getHumidity(page)
	if err != nil {
		log.Printf("Error on getting humidity value! \n%v", err)
	}
	c <- h
}

func getAirPressureAsync(page []byte, c chan uint16) {
	ap, err := getAirPressure(page)
	if err != nil {
		log.Printf("Error on getting air pressure value! \n%v", err)
	}
	c <- ap
}

func parsePageAsync(page []byte) *WeatherData {
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
