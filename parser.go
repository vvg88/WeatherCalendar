package main

import (
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	tmpExpr       = `Текущая температура<\/span><span class="temp__value">([+-]?\d{1,2})<\/span>`
	windSpeedExpr = `wind-speed">(\d{1,2}(,\d)?)<\/span>`
	weathCondExpr = `link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":\d+}}'>([а-яА-Я\s]+)`
	humExpr       = `icon_humidity-white term__fact-icon" aria-hidden="true"><\/i>(\d{1,2}|100)%`
	pressExpr     = `icon_pressure-white term__fact-icon" aria-hidden="true"><\/i>(\d{3})`
	windDirExpr   = `<abbr class=" icon-abbr" title="Ветер:\s+([а-я\-]+)"`
)

var (
	tempRegx      = regexp.MustCompile(tmpExpr)
	windSpeedRegx = regexp.MustCompile(windSpeedExpr)
	conditionRegx = regexp.MustCompile(weathCondExpr)
	humidRegx     = regexp.MustCompile(humExpr)
	pressRegx     = regexp.MustCompile(pressExpr)
	windDirRegx   = regexp.MustCompile(windDirExpr)

	errTempNotFound      = errors.New("temperature was not found")
	errWindSpeedNotFound = errors.New("wind speed was not found")
	errWeathCondNotFound = errors.New("weather condition was not found")
	errHumidityNotFound  = errors.New("humidity was not found")
	errAirPressNotFound  = errors.New("air pressure was not found")
	errWindDirNotFound   = errors.New("wind direction was not found")
)

func getTemperature(page []byte) (int, error) {
	submatches := tempRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No temperature matches found!")
		return 0, errTempNotFound
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
		return 0, errWindSpeedNotFound
	}
	wsStr := strings.ReplaceAll(string(submatches[1]), ",", ".")
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
		return "", errWeathCondNotFound
	}
	wcStr := string(submatches[1])
	log.Printf("A current weather condition is %s\n", wcStr)
	return wcStr, nil
}

func getHumidity(page []byte) (uint8, error) {
	submatches := humidRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No humidity matches found!")
		return 0, errHumidityNotFound
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
		return 0, errAirPressNotFound
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

func getWindDirection(page []byte) (string, error) {
	submatches := windDirRegx.FindSubmatch(page)
	if submatches == nil {
		log.Println("No wind direction matches found!")
		return "", errWindDirNotFound
	}
	wdStr := string(submatches[1])
	log.Printf("Wind direction: %s", wdStr)
	return wdStr, nil
}

func parsePage(name string) *WeatherData {
	page, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	t, _ := getTemperature(page)
	ws, _ := getWindSpeed(page)
	wd, _ := getWindDirection(page)
	wc, _ := getWeatherCondition(page)
	h, _ := getHumidity(page)
	ap, _ := getAirPressure(page)
	ts := time.Now()
	return &WeatherData{Temperature: t, WindSpeed: ws, WindDirection: wd, WeatherCondition: wc, Humidity: h, AirPressure: ap, TimeStamp: ts}
}
