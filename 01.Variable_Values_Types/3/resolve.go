package main

import (
	"fmt"
)

var (
	x = 42
	y = "James Bond"
	z = true
)

func main() {
	s := fmt.Sprintf("%d, %s, %t", x, y, z)
	fmt.Println(s)
}
