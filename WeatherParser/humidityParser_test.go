package main

import (
	"testing"
)

var humidityTestCases = []struct {
	name      string
	humidStr  string
	wantHumid uint8
	wantErr   error
}{
	{
		"humidGood1",
		`<i class="icon icon_humidity-white term__fact-icon" aria-hidden="true"></i>90%</div>`,
		90,
		nil,
	},
	{
		"humidGood2",
		`<i class="icon icon_humidity-white term__fact-icon" aria-hidden="true"></i>9%</div>`,
		9,
		nil,
	},
	{
		"humidGood3",
		`<i class="icon icon_humidity-white term__fact-icon" aria-hidden="true"></i>50%</div><div class="term term_orient_v fact__pressure"></i>40%`,
		50,
		nil,
	},
	{
		"humidGood4",
		`<i class="icon icon_humidity-white term__fact-icon" aria-hidden="true"></i>100%</div>`,
		100,
		nil,
	},
	{
		"humidGood5",
		`<i class="icon icon_humidity-white term__fact-icon" aria-hidden="true"></i>0%</div>`,
		0,
		nil,
	},
	{
		"humidBad1",
		`<i class="icon icon_humidity-white term__fact-icon" aria-hidden="true"></i>101%</div>`,
		0,
		errHumidityNotFound,
	},
	{
		"humidBad2",
		`<i class="icon icon_humidity-white term__fact-icon" aria-hidden="true"></i>%</div>`,
		0,
		errHumidityNotFound,
	},
	{
		"humidBad3",
		`<i class="icon icon_humidity-white term__fakt-icon" aria-hidden="true"></i>50%</div>`,
		0,
		errHumidityNotFound,
	},
}

func TestGetHumidity(t *testing.T) {
	for _, tc := range humidityTestCases {
		t.Run(tc.name, func(t *testing.T) {
			h, err := getHumidity([]byte(tc.humidStr))
			if h != tc.wantHumid {
				t.Errorf("Expected humidity: %d; actual: %d", tc.wantHumid, h)
			}
			if err != tc.wantErr {
				t.Errorf("Expected error: %v; actual: %v", tc.wantErr, err)
			}
		})
	}
}
