package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

const projectID = "weather-calend-fires"

func createClient(ctx context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	// Close client when done with defer client.Close()
	return client, nil
}
