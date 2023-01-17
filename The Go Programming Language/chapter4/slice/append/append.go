package main

import "fmt"

func appendInt(s []int, i int) []int {
	if len(s) < cap(s) {
		// panic: runtime error: index out of range [3] with length 3
		// s[len(s)] = i
		z := s[:len(s)+1]
		z[len(s)] = i

		return z
	} else {
		c := 2 * cap(s)
		d := make([]int, c)

		n := copy(d, s)
		fmt.Println("copy ", n)

		d[cap(s)] = i

		return d
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("a: %d %d %p, %p, %p, %v\n", len(a), cap(a), &a, a, &a[0], a)

	b := a[:3]
	fmt.Printf("b: %d %d %p, %p, %p, %v\n", len(b), cap(b), &b, b, &b[0], b)

	c := appendInt(b, 6)
	fmt.Printf("c: %d %d %p, %p, %p, %v\n", len(c), cap(c), &c, c, &c[0], c)

	d := appendInt(c, 7)
	fmt.Printf("d: %d %d %p, %p, %p, %v\n", len(d), cap(d), &d, d, &d[0], d)

	e := appendInt(d, 8)
	fmt.Printf("e: %d %d %p, %p, %p, %v\n", len(e), cap(e), &e, e, &e[0], e)
}
