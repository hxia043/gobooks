package main

import (
	"strings"
)

func plusContact(n int, testString string) {
	s, sep := "", ""

	for i := 0; i < n; i++ {
		s += sep + testString
		sep = "-"
	}
}

func stringsJoinContact(n int, testString string) {
	s := make([]string, 0)

	for i := 0; i < n; i++ {
		s = append(s, testString)
	}

	sep := "-"
	strings.Join(s, sep)
}

func stringsJoinContactPre(n int, testString string) {
	s := make([]string, n)

	for i := 0; i < n; i++ {
		s = append(s, testString)
	}

	sep := "-"
	strings.Join(s, sep)
}

func main() {
	plusContact(3, "helloworld")
	stringsJoinContact(3, "helloworld")
	stringsJoinContactPre(3, "helloworld")
}
