package main

import (
	"fmt"
)

func main() {
	s := []byte{104, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100}

	fmt.Println(s, len(s), cap(s))

	s = append(s, 0)[:len(s)]

	fmt.Println(s, len(s), cap(s))
}
