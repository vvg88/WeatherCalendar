package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// WeatherData contains weather condition data
type WeatherData struct {
	Temperature      int       `json:"temp"`
	WindSpeed        float32   `json:"windSpd"`
	WindDirection    string    `json:"windDir"`
	WeatherCondition string    `json:"weatherCond"`
	Humidity         uint8     `json:"humidity"`
	AirPressure      uint16    `json:"airPress"`
	TimeStamp        time.Time `json:"timeStamp"`
}

func (wd *WeatherData) save() error {
	b, err := wd.toJSON()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("weathDat.json", b, 0600)
	if err != nil {
		log.Printf("Error on writing serialized weather condition data: %v\n", err)
		return err
	}
	return nil
}

func (wd *WeatherData) toJSON() (data []byte, err error) {
	data, err = json.Marshal(wd)
	if err != nil {
		log.Printf("Error on marshaling weather data: %+v; error: %v\n", wd, err)
		return nil, err
	}
	return data, nil
}
