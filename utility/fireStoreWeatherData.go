package main

import (
	"fmt"
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
	key := fswd.key()
	_, err := fsClient.Collection(weatherDataCollectionName).Doc(key).Set(ctx, fswd)
	return err
}

func (fswd *fsWeatherData) key() string {
	ts := fswd.TimeStamp
	return fmt.Sprintf("%d.%d.%d-%d", ts.Day(), ts.Month(), ts.Year(), ts.Hour())
}

func (fswd *fsWeatherData) setLocalTZ() error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}
	fswd.TimeStamp = fswd.TimeStamp.In(loc)
	return nil
}
