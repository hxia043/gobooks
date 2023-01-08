package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string, result chan string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		result <- fmt.Sprintf("reading %s %v\n", url, err)
		return
	}

	//n, err := io.Copy(io.Discard, resp.Body)
	file, _ := os.Create("output.html")
	n, err := io.Copy(file, resp.Body)
	if err != nil {
		result <- fmt.Sprintf("copy %v\n", err)
		return
	}

	resp.Body.Close()

	secs := time.Since(start).Seconds()
	result <- fmt.Sprintf("%s: reading %s %d in %.2f seconds\n", resp.Status, url, n, secs)
}

func main() {
	start := time.Now()
	result := make(chan string)

	for _, arg := range os.Args[1:] {
		go fetch(arg, result)
	}

	/*
		for r := range result {
			fmt.Println(r)
		}
	*/

	for range os.Args[1:] {
		fmt.Println(<-result)
	}

	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}
