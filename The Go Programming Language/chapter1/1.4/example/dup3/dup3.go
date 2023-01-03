package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./file")
	if err != nil {
		log.Fatalln(err)
	}

	counts := make(map[string]int)
	/*
		s := strings.Split(string(data), "\n")
		for _, line := range s {
			counts[line]++
		}
	*/

	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
