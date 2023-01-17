package main

import "fmt"

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}

	return s
}

func main() {
	a := []rune{1, 2, 3, 4, 5}
	fmt.Printf("%T %v %p %p %p %p %p\n", a, a, &a, a, &a[0], &a[1], &a[2])

	b := reverse(a)
	fmt.Printf("%T %v %p %p %p %p %p\n", b, b, &b, b, &b[0], &b[1], &b[2])

	// shift n elements
	// n = 2
	a = reverse(a)
	fmt.Printf("%T %v %p %p %p %p %p\n", a, a, &a, a, &a[0], &a[1], &a[2])

	reverse(a[:2])
	fmt.Printf("%T %v %p %p %p %p %p\n", a, a, &a, a, &a[0], &a[1], &a[2])

	reverse(a)
	reverse(a[:len(a)-2])

	fmt.Printf("%T %v %p %p %p %p %p\n", a, a, &a, a, &a[0], &a[1], &a[2])
}
