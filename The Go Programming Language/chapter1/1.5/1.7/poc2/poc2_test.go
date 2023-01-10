package main

import "testing"

func BenchmarkGetFileFromReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFileFromReadAll()
	}
}

func BenchmarkGetFileFromCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFileFromReadAll()
	}
}
