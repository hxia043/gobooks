package main

import "fmt"

func main() {
	// array
	a := [3]int{0, 1, 2}
	fmt.Printf("%T, %d, %p, %p, %p\n", a, a[0], &a, &a[0], &a[1])

	b := a
	fmt.Printf("%T, %d, %p, %p, %p\n", b, b[0], &b, &b[0], &b[1])

	// nil and empty slice
	c := []rune{}
	if c == nil {
		fmt.Printf("nil slice %T %v\n", c, c)
	} else {
		fmt.Printf("empty slice %T %v %d\n", c, c, len(c))
	}

	var d []rune
	if d == nil {
		fmt.Printf("nil slice %T %v %d\n", d, d, len(d))
	} else {
		fmt.Printf("empty slice %T %v %d\n", d, d, len(d))
	}
	// how the impact of the diffenency between nil and empty slice

	// address of slice
	f := []rune{}
	fmt.Printf("%T %v %p %p\n", f, f, &f, f)

	/*
		// panic: runtime error: index out of range [0] with length 0
		f[0] = 1
		fmt.Printf("%T %v %p %p\n", f, f, &f, f)
	*/

	g := []rune{1, 2, 3, 4, 5}
	fmt.Printf("%T %v %p %p %p %p %p\n", g, g, &g, g, &g[0], &g[1], &g[2])

	h := g
	fmt.Printf("%T %v %p %p %p %p %p\n", h, h, &h, h, &h[0], &h[1], &h[2])

	i := h[1:4]
	fmt.Printf("%T %v %p %p %p %p %p\n", i, i, &i, i, &i[0], &i[1], &i[2])

	fmt.Printf("%d, %d\n", len(i), cap(i))

	j := append(i, 6)
	fmt.Println(g)
	fmt.Printf("i: %T %v %p %p %p %p %p\n", i, i, &i, i, &i[0], &i[1], &i[2])
	fmt.Printf("i len: %d %d\n", len(i), cap(i))
	fmt.Printf("j: %T %v %p %p %p %p %p\n", j, j, &j, j, &j[0], &j[1], &j[2])

	i = append(i, 7)
	fmt.Printf("i: %T %v %p %p %p %p %p\n", i, i, &i, i, &i[0], &i[1], &i[2])
	fmt.Printf("i len: %d %d\n", len(i), cap(i))

	i = append(i, 8)
	fmt.Printf("i: %T %v %p %p %p %p %p\n", i, i, &i, i, &i[0], &i[1], &i[2])
	fmt.Printf("i len: %d %d\n", len(i), cap(i))
}
