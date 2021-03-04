package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const projectID = "weather-calend-fires"

var fsClient *firestore.Client
var ctx context.Context

func main() {
	ctx = context.Background()
	var err error
	fsClient, err = firestore.NewClient(ctx, projectID)
	defer fsClient.Close()
	if err != nil {
		log.Fatalf("Unable to create firestore client!\nError: %v", err)
	}

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
