package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	fmt.Println(unsafe.Sizeof(a))

	var b int8 = 127
	fmt.Println(unsafe.Sizeof(b), b+1)
	// cannot use 65535 (untyped int constant) as int8 value in assignment (overflows)
	// b = 65535

	var c uint8
	c = 255
	fmt.Println(unsafe.Sizeof(c), c+1)

	var d float32 = 16777216
	fmt.Println(unsafe.Sizeof(d), d == d+1)

	var e float64 = 16777216
	fmt.Println(unsafe.Sizeof(e), e == e+1)

	var f byte
	f = c
	fmt.Println(unsafe.Sizeof(f), f)

	var g = 0
	var h = int8(0)
	fmt.Printf("%T, %d, %d\n", g, unsafe.Sizeof(g), g)
	fmt.Printf("%T, %d, %d\n", h, unsafe.Sizeof(h), h)

	var i = "严"
	fmt.Println(unsafe.Sizeof(i), i)
}
