package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func mainTest() {
	/*
		file, _ := os.Open("./output")
		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprint(os.Stdout, err)
		}

		fmt.Println(string(data))
	*/
	fmt.Println("test")
}

func main() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		file, err := os.Create("output")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		/*
			d1 := make([]byte, 1023)
			resp.Body.Read(d1)
			fmt.Println(string(d1))
		*/

		if _, err := io.Copy(file, resp.Body); err != nil {
			fmt.Fprint(os.Stderr, err)
		}

		resp.Body.Close()

		/*
			d2 := make([]byte, 1024)
			resp.Body.Read(d2)
			fmt.Println(string(d2))
		*/

		/*
			data, err := io.ReadAll(file)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
			}

			fmt.Printf("data %s\n", data)
		*/
	}

	mainTest()
}
