package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetFileFromCopy() {
	//start := time.Now()

	resp, _ := http.Get("https://artifactory-hz1.int.net.nokia.com/artifactory/mnp5g-central-public-remote/System_Release/vDUCNF00/vDUCNF00_0.300.10253/artifacts/log-collector-v2.3.tar.gz")
	file, _ := os.Create("log-collector-v2.3.tar.gz")

	//t1 := time.Now()
	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(time.Since(t1).Seconds())

	//fmt.Println(time.Since(start).Seconds())
}

func GetFileFromReadAll() {
	//start := time.Now()

	resp, _ := http.Get("https://artifactory-hz1.int.net.nokia.com/artifactory/mnp5g-central-public-remote/System_Release/vDUCNF00/vDUCNF00_0.300.10253/artifacts/log-collector-v2.3.tar.gz")

	//t1 := time.Now()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile("./log-collector-v2.3.tar.gz", data, 0755); err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(time.Since(t1).Seconds())

	//fmt.Println(time.Since(start).Seconds())
}

func main() {
	GetFileFromCopy()
	GetFileFromReadAll()
}
