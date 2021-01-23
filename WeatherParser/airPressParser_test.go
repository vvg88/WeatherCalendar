package main

import (
	"testing"
)

var apTestCases = []struct {
	name     string
	apStr    string
	airPress uint16
	wantErr  error
}{
	{
		"airPressGood1",
		`<i class="icon icon_pressure-white term__fact-icon" aria-hidden="true"></i>750 <span class="fact__unit">мм рт. ст.</span>`,
		750,
		nil,
	},
	{
		"airPressGood2",
		`<i class="icon icon_pressure-white term__fact-icon" aria-hidden="true"></i>750 <span class="fact__unit">мм рт. ст.</span><span class="fact__unit" aria-hidden="true"></i>760 `,
		750,
		nil,
	},
	{
		"airPressBad1",
		`<i class="icon icon_pressure-white term__fact-icon" aria-hidden="true"></i>73 <span class="fact__unit">мм рт. ст.</span>`,
		0,
		errAirPressNotFound,
	},
	{
		"airPressBad2",
		`<i class="icon icon_pressure-white term__fact-icon" aria-hidden="true"></i>-750 <span class="fact__unit">мм рт. ст.</span>`,
		0,
		errAirPressNotFound,
	},
	{
		"airPressBad3",
		`<i class="icon icon_pressure-white term__fact-icon" aria-hiden="true"></i>750 <span class="fact__unit">мм рт. ст.</span>`,
		0,
		errAirPressNotFound,
	},
}

func TestGetAirPress(t *testing.T) {
	for _, tc := range apTestCases {
		t.Run(tc.name, func(t *testing.T) {
			ap, err := getAirPressure([]byte(tc.apStr))
			if ap != tc.airPress {
				t.Errorf("Expected air pressure: %d; actual: %d", tc.airPress, ap)
			}
			if err != tc.wantErr {
				t.Errorf("Expected error: %v; actual: %v", tc.wantErr, err)
			}
		})
	}
}
