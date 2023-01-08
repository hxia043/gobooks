package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: getting %s: %v\n", resp.Status, arg, err)
		}

		fmt.Fprintf(os.Stdout, "%s: getting %s finised\n", resp.Status, arg)

		file, err := os.Create("outputhtml")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		if _, err := io.Copy(file, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
		}

		resp.Body.Close()
	}
}
