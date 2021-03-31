package WeatherCalendarCommon

import "time"

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
