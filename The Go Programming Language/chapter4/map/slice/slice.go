package main

import "fmt"

func compare(s1, s2 []int) bool {
	if s1 == nil || s2 == nil {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	for i, j := 0, 0; i < len(s1); i++ {
		if s1[i] == s2[j] {
			j++
		} else {
			return false
		}
	}

	return true
}

func main() {
	//s1 := []int{1, 2, 3, 4, 5}
	//s2 := []int{1, 2, 3, 4, 5}

	//s1 := make([]int, 0)
	//s2 := make([]int, 0)

	var s1 []int
	var s2 []int
	if compare(s1, s2) {
		fmt.Println("same slice")
	}
}
