package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countLines(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}

	isRepeat := false
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			isRepeat = true
		}
	}

	return isRepeat
}

func main() {
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			log.Fatalln(err)
		}

		defer func() {
			if err := f.Close(); err != nil {
				log.Fatalln(err)
			}
		}()

		if countLines(f) {
			fmt.Println("Info: repeat line file is ", arg)
		}
	}
}
