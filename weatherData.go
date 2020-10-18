package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// WeatherData contains weather condition data
type WeatherData struct {
	Temperature      int
	WindSpeed        float32
	WeatherCondition string
	Humidity         uint8
	AirPressure      uint16
}

func (wd *WeatherData) save() error {
	b, err := json.Marshal(wd)
	if err != nil {
		log.Printf("Error gotten on marshaling weather data: %+v; error: %v\n", wd, err)
		return err
	}
	err = ioutil.WriteFile("weathDat.json", b, 0600)
	if err != nil {
		log.Printf("Error on writing serialized weather condition data: %v\n", err)
		return err
	}
	return nil
}
