package main

import (
	"testing"
)

const testYaPageFile = "yaPageTest.txt"

func TestParserFull(t *testing.T) {
	page, err := readHTTPPage()
	if err != nil {
		t.Errorf("Read Weather page failed with error: %v", err)
	}
	err = savePage(testYaPageFile, page)
	if err != nil {
		t.Errorf("Save Weather page failed with error: %v", err)
	}
	page, err = openPage(testYaPageFile)
	if err != nil {
		t.Errorf("Open Weather page failed with error: %v", err)
	}
	wd, err := parsePageSync(page)
	if err != nil {
		t.Errorf("Parse Weather page failed with error: %v", err)
	}
	err = wd.save()
	if err != nil {
		t.Errorf("Save parsed Weather page failed with error: %v", err)
	}
}

func TestYaPageAvailable(t *testing.T) {
	_, err := readHTTPPage()
	if err != nil {
		t.Errorf("Weather page reading failed with error: %v", err)
	}
}

func BenchmarkSyncParser(b *testing.B) {
	page, _ := openPage("yaPage.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parsePageSync(page)
	}
}

func BenchmarkAsyncParser(b *testing.B) {
	page, _ := openPage("yaPage.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parsePageAsync(page)
	}
}
