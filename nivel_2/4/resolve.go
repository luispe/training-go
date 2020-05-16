package main

import (
	"fmt"
)

func main() {
	a := 11
	fmt.Printf("%d\n%b\n%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\n%b\n%#x", b, b, b)
}
