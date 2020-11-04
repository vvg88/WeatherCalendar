package main

import (
	"testing"
)

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
