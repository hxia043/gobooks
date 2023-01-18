package main

import "fmt"

func main() {
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%v %p %p %p\n", a1, &a1, &a1[0], &a1[2])
	fmt.Printf("%v %p %p %p\n", a2, &a2, &a2[0], &a2[2])

	if a1 == a2 {
		fmt.Println("same array")
	}
}
