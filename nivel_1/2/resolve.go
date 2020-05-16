package main

import (
	"fmt"
)

var (
	x int
	y string
	z bool
)

// 2-b. Variables zero (0) value, cuando se asignan en memoria se crean con lo que se conoce como zero values, esto
// significa que para un tipo string sera una cadena vacía "" para int 0, bool false, etc.
// + info https://golang.org/ref/spec#The_zero_value

func main() {
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
