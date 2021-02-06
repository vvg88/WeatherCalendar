package main

import "time"

type weatherData struct {
	Temperature      int       `json:"temp"`
	WindSpeed        float32   `json:"windSpd"`
	WindDirection    string    `json:"windDir"`
	WeatherCondition string    `json:"weatherCond"`
	Humidity         uint8     `json:"humidity"`
	AirPressure      uint16    `json:"airPress"`
	TimeStamp        time.Time `json:"timeStamp"`
}

func (wd *weatherData) toFirestoreData() *fsWeatherData {
	return &fsWeatherData{
		Temperature:      wd.Temperature,
		WindSpeed:        wd.WindSpeed,
		WindDirection:    wd.WindDirection,
		WeatherCondition: wd.WeatherCondition,
		Humidity:         wd.Humidity,
		AirPressure:      wd.AirPressure,
		TimeStamp:        wd.TimeStamp,
	}
}
