package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	a bool
	b int32
	c string
	d string
}

type s2 struct {
	b int32
	c string
	d string
	a bool
}

func main() {
	var s1 s1
	fmt.Println(unsafe.Sizeof(s1))

	var s2 s2
	fmt.Println(unsafe.Sizeof(s2))
}
