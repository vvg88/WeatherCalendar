package main

import "testing"

func TestYaPageAvailable(t *testing.T) {
	_, err := readHTTPPage()
	if err != nil {
		t.Errorf("Reading of weather page failed with error: %v", err)
	}
}
