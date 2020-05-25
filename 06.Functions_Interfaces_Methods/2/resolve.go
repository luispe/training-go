package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Println(foo(s1...))

	s2 := []int{1, 2, 3, 4}
	fmt.Println(bar(s2))
}

func foo(xs ...int) int {
	total := 0
	for _, v := range xs {
		total = total + v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total = total + v
	}
	return total
}
