package main

import (
	"testing"
)

var getWsTestCases = []struct {
	name      string
	wsPageStr string
	wantVal   float32
	wantErr   error
}{
	{
		"wsGood1",
		`<span class="wind-speed">5</span> <span class="fact__unit">`,
		5,
		nil,
	},
	{
		"wsGood2",
		`<span class="wind-speed">15</span> <span class="fact__unit">`,
		15,
		nil,
	},
	{
		"wsGood3",
		`<span class="wind-speed">5,2</span> <span class="fact__unit">`,
		5.2,
		nil,
	},
	{
		"wsGood4",
		`<span class="wind-speed">15,2</span> <span class="fact__unit">`,
		15.2,
		nil,
	},
	{
		"wsGood5",
		`<span class="wind-speed">5,5</span> <span class="fact__unit">15,5</span>`,
		5.5,
		nil,
	},
	{
		"wsBad1",
		`<span class="wind-speed">152</span> <span class="fact__unit">`,
		0,
		errWindSpeedNotFound,
	},
	{
		"wsBad2",
		`<span class="wind-speed">152,3</span> <span class="fact__unit">`,
		0,
		errWindSpeedNotFound,
	},
	{
		"wsBad3",
		`<span class="wind-speed">1,35</span> <span class="fact__unit">`,
		0,
		errWindSpeedNotFound,
	},
	{
		"wsBad4",
		`<span class="wind-speed">5.5</span> <span class="fact__unit">`,
		0,
		errWindSpeedNotFound,
	},
}

func TestGetWindSpeed(t *testing.T) {
	for _, tc := range getWsTestCases {
		t.Run(tc.name, func(t *testing.T) {
			ws, err := getWindSpeed([]byte(tc.wsPageStr))
			if ws != tc.wantVal {
				t.Errorf("Expected wind speed: %.1f; actual: %.1f", tc.wantVal, ws)
			}
			if err != tc.wantErr {
				t.Errorf("Expected error: %v; actual: %v", tc.wantErr, err)
			}
		})
	}
}
