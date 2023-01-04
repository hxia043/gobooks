package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countLines(f *os.File) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	file, err := os.Open("./file")
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err != nil {
		log.Fatalln(err)
	}

	countLines(file)
}
