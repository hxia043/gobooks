package main

import (
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func benchmark(b *testing.B, f func(int, string)) {
	for i := 0; i < b.N; i++ {
		f(30, randomString(10))
	}
}

func BenchmarkPlusContact(b *testing.B)           { benchmark(b, plusContact) }
func BenchmarkStringsJoinContactPre(b *testing.B) { benchmark(b, stringsJoinContactPre) }
func BenchmarkStringsJoinContact(b *testing.B)    { benchmark(b, stringsJoinContact) }
