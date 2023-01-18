package main

import "fmt"

func compareString(m1, m2 map[string]string) bool {
	if m1 == nil || m2 == nil {
		return false
	}

	if len(m1) == 0 && len(m2) == 0 {
		return true
	}

	if len(m1) != len(m2) {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok {
			return false
		} else {
			if v1 != v2 {
				return false
			}
		}
	}

	return true
}

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

func compareSlice(m1, m2 map[string][]int) bool {
	if m1 == nil || m2 == nil {
		return false
	}

	// make empty are not compared
	if len(m1) == 0 || len(m2) == 0 {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			if !compare(v1, v2) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	//var m1 map[string]string
	//var m2 map[string]string

	// invalid operation: cannot compare m1 == m2 (map can only be compared to nil)
	// if m1 == m2 {}

	//m1 := make(map[string]string)
	//m2 := make(map[string]string)

	/*
		m1 := map[string]string{
			"name":    "hxia",
			"company": "Nokia",
		}

		m2 := map[string]string{
			"name":    "hxia",
			"company": "Nokia",
		}

		if compareString(m1, m2) {
			fmt.Println("same map")
		}
	*/

	s1 := map[string][]int{
		"age": {1, 2, 3, 4, 5},
	}

	s2 := map[string][]int{
		"abe": {1, 2, 3, 4, 5},
	}

	if compareSlice(s1, s2) {
		fmt.Println("same map slice")
	}
}
