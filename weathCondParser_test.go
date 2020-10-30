package main

import (
	"testing"
)

var testCases = []struct {
	name         string
	weathCondStr string
	want         string
	wantErr      error
}{
	{
		"wcGood1",
		`<div class="link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":18}}'>Небольшой дождь</div>`,
		"Небольшой дождь",
		nil,
	},
	{
		"wcGood2",
		`<div class="link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":18}}'>Ясно</div>`,
		"Ясно",
		nil,
	},
	{
		"wcGood3",
		`<div class="link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":18}}'>Облачно с прояснениями</div>`,
		"Облачно с прояснениями",
		nil,
	},
	{
		"wcGood4",
		`<div class="link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":1}}'>Ясно</div><div class="term">Облачно</div>`,
		"Ясно",
		nil,
	},
	{
		"wcBad1",
		`<div class="link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":1}}'>Condition</div>`,
		"",
		errWeathCondNotFound,
	},
	{
		"wcBad2",
		`<div class="link__condition day-anchor i-bem" data-bem='{"day-ancсhor":{"anchor":1}}'>Ясно</div>`,
		"",
		errWeathCondNotFound,
	},
	{
		"wcBad3",
		`<div class="link__condition day-anchor i-bem" data-bem='{"day-anchor":{"anchor":}}'>Ясно</div>`,
		"",
		errWeathCondNotFound,
	},
}

func TestGetWeatherCondition(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wc, err := getWeatherCondition([]byte(tc.weathCondStr))
			if wc != tc.want {
				t.Errorf("Expected weather condition: %s; actual: %s", tc.want, wc)
			}
			if err != tc.wantErr {
				t.Errorf("Expected error: %v; actual: %v", tc.wantErr, err)
			}
		})
	}
}
