package main

import "fmt"

type stack []int

func (s stack) push(d int) stack {
	if s == nil {
		fmt.Println("stack is a nil")
		return nil
	}

	if len(s) < cap(s) {
		s = append(s, d)
		fmt.Println(s)
		return s
	}

	// for easy demo make stack as a limit top
	fmt.Println("trigger stack top limit, do nothing")
	return s
}

func (s stack) pop() stack {
	if s == nil {
		fmt.Println("stack is a nil")
		return nil
	}

	if len(s) == 0 {
		fmt.Println("trigger stack bottom limit, do nothing")
		return s
	}

	s = s[:len(s)-1]
	fmt.Println(s)

	return s
}

func main() {
	s := make(stack, 24)[:0]
	fmt.Printf("s: %v len: %d cap: %d\n", s, len(s), cap(s))

	s = s.push(1)
	s = s.push(2)
	s = s.push(3)
	s = s.pop()
	s = s.pop()
	s = s.pop()
}
