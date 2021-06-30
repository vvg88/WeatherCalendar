package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const projectID = "weather-calend-fires"

var mode = flag.String("mode", "", "the command the utility should run.")

var fsClient *firestore.Client
var ctx context.Context

func main() {
	flag.Parse()

	ctx = context.Background()
	err := initFsClient()
	defer fsClient.Close()
	if err != nil {
		log.Fatalf("Unable to create firestore client!\nError: %v", err)
	}

	switch *mode {
	case "backupData":
		backUpData()
	case "restoreData":
		restoreData()
	default:
		fmt.Println("Specify the mode utility should run.")
	}
}

func initFsClient() error {
	var err error
	fsClient, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	return nil
}

func backUpData() {
	iter := fsClient.Collection(weatherDataCollectionName).Documents(ctx)
	var wdSlice []fsWeatherData
	for cnt := 0; ; cnt++ {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("Error on getting next doc snap! Error: %v\n", err)
			break
		}

		// Get document snap
		var wd fsWeatherData
		err = docSnap.DataTo(&wd)
		if err != nil {
			fmt.Printf("Error on converting firestore record! Error: %v", err)
			continue
		}
		err = wd.setLocalTZ()
		if err != nil {
			fmt.Printf("Unable to load local time zone! Error: %v", err)
		}
		wdSlice = append(wdSlice, wd)
		fmt.Printf("%d:\t%+v\n", cnt, wd)
	}
	//fmt.Println(wdSlice)

	data, err := json.Marshal(wdSlice)
	if err != nil {
		fmt.Printf("Error on marshaling weather data! error: %v\n", err)
	}

	err = ioutil.WriteFile("backupWeatherData.json", data, 0600)
	if err != nil {
		fmt.Printf("Error on writing serialized weather data: %v\n", err)
	}
}

func renameKeys() {
	iter := fsClient.Collection(weatherDataCollectionName).Documents(ctx)
	var wd fsWeatherData
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}

		// Get document snap
		err = docSnap.DataTo(&wd)
		if err != nil {
			log.Println(err)
			continue
		}

		// Continue if the doc ID is new
		/*if strings.Contains(docID, ".") {
			continue
		}
		fmt.Printf("Time stamp: %v\n", wd.TimeStamp)
		fmt.Printf("Time stamp biased: %v\n", wd.TimeStamp)

		break*/

		// Save the doc with new ID and biased time stamp
		loc, err := time.LoadLocation("Local")
		wd.TimeStamp = wd.TimeStamp.In(loc)
		err = wd.save()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("Doc '%s' saved!\n", wd.key())

		// Delete the doc with old ID
		docID := docSnap.Ref.ID
		_, err = fsClient.Collection(weatherDataCollectionName).Doc(docID).Delete(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("Doc '%s' deleted!\n\n", docID)
	}

	fmt.Println("Done!")
}

func restoreData() {
	loc, err := time.LoadLocation("Local")
	timeStamp := time.Date(2021, time.April, 8, 1, 0, 0, 0, loc)
	fsWd := fsWeatherData{
		TimeStamp:        timeStamp,
		Temperature:      3,
		AirPressure:      739,
		Humidity:         92,
		WeatherCondition: "Пасмурно",
		WindDirection:    "восточный",
		WindSpeed:        4,
	}
	err = fsWd.save()
	if err != nil {
		fmt.Printf("Error on saving lost values: %+v", err)
	}
	fmt.Println("Lost data successfully saved!")

	timeStamp = time.Date(2021, time.April, 8, 7, 0, 0, 0, loc)
	fsWd = fsWeatherData{
		TimeStamp:        timeStamp,
		Temperature:      4,
		AirPressure:      734,
		Humidity:         96,
		WeatherCondition: "Дождь",
		WindDirection:    "южный",
		WindSpeed:        4,
	}
	err = fsWd.save()
	if err != nil {
		fmt.Printf("Error on saving lost values: %+v", err)
	}
	fmt.Println("Lost data successfully saved!")

	timeStamp = time.Date(2021, time.April, 8, 13, 0, 0, 0, loc)
	fsWd = fsWeatherData{
		TimeStamp:        timeStamp,
		Temperature:      5,
		AirPressure:      734,
		Humidity:         92,
		WeatherCondition: "Дождь",
		WindDirection:    "южный",
		WindSpeed:        6,
	}
	err = fsWd.save()
	if err != nil {
		fmt.Printf("Error on saving lost values: %+v", err)
	}
	fmt.Println("Lost data successfully saved!")

	timeStamp = time.Date(2021, time.April, 8, 19, 0, 0, 0, loc)
	fsWd = fsWeatherData{
		TimeStamp:        timeStamp,
		Temperature:      1,
		AirPressure:      740,
		Humidity:         92,
		WeatherCondition: "Слабый снег",
		WindDirection:    "юго-западный",
		WindSpeed:        6,
	}
	err = fsWd.save()
	if err != nil {
		fmt.Printf("Error on saving lost values: %+v", err)
	}
	fmt.Println("Lost data successfully saved!")
}
