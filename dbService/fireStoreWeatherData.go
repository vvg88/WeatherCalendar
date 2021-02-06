package main

import (
	"time"
)

const weatherDataCollectionName = "weather-data"

type fsWeatherData struct {
	Temperature      int       `firestore:"Temperature,omitempty"`
	WindSpeed        float32   `firestore:"WindSpeed,omitempty"`
	WindDirection    string    `firestore:"WindDirection,omitempty"`
	WeatherCondition string    `firestore:"WeatherCondition,omitempty"`
	Humidity         uint8     `firestore:"Humidity,omitempty"`
	AirPressure      uint16    `firestore:"AirPressure,omitempty"`
	TimeStamp        time.Time `firestore:"TimeStamp,omitempty"`
}

func (fswd *fsWeatherData) save() error {
	_, _, err := fsClient.Collection(weatherDataCollectionName).Add(ctx, fswd)
	return err
}
