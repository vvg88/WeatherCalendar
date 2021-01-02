package main

import "testing"

var getTempTestCases = []struct {
	name        string
	tempPageStr string
	wantVal     int
	wantErr     error
}{
	{
		name:        "tempGood1",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">+3</span>`,
		wantVal:     3,
		wantErr:     nil,
	},
	{
		name:        "tempGood2",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">+26</span>`,
		wantVal:     26,
		wantErr:     nil,
	},
	{
		name:        "tempGood3",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">−3</span>`,
		wantVal:     -3,
		wantErr:     nil,
	},
	{
		name:        "tempGood4",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">−25</span>`,
		wantVal:     -25,
		wantErr:     nil,
	},
	{
		name:        "tempGood5",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">0</span>`,
		wantVal:     0,
		wantErr:     nil,
	},
	{
		name:        "tempBad1",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">100</span>`,
		wantVal:     0,
		wantErr:     errTempNotFound,
	},
	{
		name:        "tempBad2",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">--1</span>`,
		wantVal:     0,
		wantErr:     errTempNotFound,
	},
	{
		name:        "tempBad3",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Random text</span><span class="temp__value">100</span>`,
		wantVal:     0,
		wantErr:     errTempNotFound,
	},
	{
		name:        "tempBad4",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</spann><span class="temp__value">10</span>`,
		wantVal:     0,
		wantErr:     errTempNotFound,
	},
	{
		name:        "tempBad5",
		tempPageStr: `<span class="temp__pre-a11y a11y-hidden">Текущая температура</span><span class="temp__value">-5</span>`,
		wantVal:     0,
		wantErr:     errTempNotFound,
	},
}

func TestGetTemperature(t *testing.T) {
	for _, tc := range getTempTestCases {
		t.Run(tc.name, func(t *testing.T) {
			tmp, err := getTemperature([]byte(tc.tempPageStr))
			if err != tc.wantErr {
				t.Errorf("Get temperature returned an unexpected error! err: %+v; expected: %+v", err, tc.wantErr)
			}
			if tc.wantVal != tmp {
				t.Errorf("Wrong temperature value! Want: %d; gotten: %d", tc.wantVal, tmp)
			}
		})
	}
}
