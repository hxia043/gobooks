package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		url := ""
		if !(strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "https://")) {
			url = fmt.Sprintf("%s%s", "http://", arg)
		} else {
			url = arg
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		file, err := os.Create("output")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		if _, err := io.Copy(file, resp.Body); err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		resp.Body.Close()
	}
}
