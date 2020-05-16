package main

import (
	"fmt"
)

const (
	a = 2020 + iota
	b = 2020 + iota
	c = 2020 + iota
	d = 2020 + iota
)

func main() {
	fmt.Printf("%d\n%d\n%d\n%d", a, b, c, d)
}
