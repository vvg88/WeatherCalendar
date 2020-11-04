package main

import (
	"testing"
)

var wdTestCases = []struct {
	name    string
	wdStr   string
	windDir string
	wantErr error
}{
	{
		"wdGood1",
		` <abbr class=" icon-abbr" title="Ветер: юго-западный" role="text">ЮЗ</abbr>`,
		"юго-западный",
		nil,
	},
	{
		"wdGood2",
		` <abbr class=" icon-abbr" title="Ветер: западный" aria-label="Ветер: юго-западный" role="text">ЮЗ</abbr>`,
		"западный",
		nil,
	},
	{
		"wdBad1",
		` <abbr class=" icon-abbr" title="Ветер: юго западный" aria-label="Ветер: юго-западный" role="text">ЮЗ</abbr>`,
		"",
		errWindDirNotFound,
	},
	{
		"wdBad2",
		` <abbr class=" icon-abbr" title="Ветер: Западный" aria-label="Ветер: юго-западный" role="text">ЮЗ</abbr>`,
		"",
		errWindDirNotFound,
	},
	{
		"wdBad3",
		` <abbr class=" icon-abr" title="Ветер: западный" aria-label="Ветер: юго-западный" role="text">ЮЗ</abbr>`,
		"",
		errWindDirNotFound,
	},
	{
		"wdBad4",
		` <abbr class=" icon-abbr" title="Ветер:западный" aria-label="Ветер: юго-западный" role="text">ЮЗ</abbr>`,
		"",
		errWindDirNotFound,
	},
}

func TestGetWindDir(t *testing.T) {
	for _, tc := range wdTestCases {
		t.Run(tc.name, func(t *testing.T) {
			wd, err := getWindDirection([]byte(tc.wdStr))
			if wd != tc.windDir {
				t.Errorf("Expected wind direction: %s; actual: %s", tc.windDir, wd)
			}
			if err != tc.wantErr {
				t.Errorf("Expected error: %v; actual: %v", tc.wantErr, err)
			}
		})
	}
}
