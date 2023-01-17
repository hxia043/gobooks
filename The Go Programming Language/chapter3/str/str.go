package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a string
	a = "hello, world"
	fmt.Println(unsafe.Sizeof(a), a)

	b := a
	fmt.Println(unsafe.Sizeof(b), b)

	a = "hello, hxia"
	fmt.Println(unsafe.Sizeof(a), a)

	fmt.Println(unsafe.Sizeof(b), b)
}
