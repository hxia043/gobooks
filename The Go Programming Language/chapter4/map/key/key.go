package main

import "fmt"

func k(s []string) string {
	return fmt.Sprintf("%q", s)
}

func main() {
	m := map[string]int{}

	s1 := []string{"hxia", "29", "male"}
	s2 := []string{"huyun", "29", "male"}

	m[k(s1)]++
	m[k(s2)]++

	fmt.Println("m: ", m)

	m[k(s1)]++
	fmt.Println("m: ", m)
}
