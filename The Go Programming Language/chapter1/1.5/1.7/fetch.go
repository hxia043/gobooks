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
			fmt.Println("1")
			fmt.Fprint(os.Stderr, err)
		}

		//data, err := io.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprint(os.Stderr, err)
		//}

		// fmt.Printf("%s", data)

		file, err := os.Create("output")
		if err != nil {
			fmt.Println("2")
			fmt.Fprint(os.Stderr, err)
		}

		//os.NewFile()
		if _, err := io.Copy(file, resp.Body); err != nil {
			fmt.Println("3")
			fmt.Fprint(os.Stderr, err)
		}

		resp.Body.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("4")
			fmt.Fprint(os.Stderr, err)
		}

		fmt.Printf("data %s", data)
	}
}
