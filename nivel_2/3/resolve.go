package main

import (
	"fmt"
)

var (
	x     = 42
	y int = 11
)

func main() {
	s := fmt.Sprintf("%d, %d", x, y)
	fmt.Println(s)
}
