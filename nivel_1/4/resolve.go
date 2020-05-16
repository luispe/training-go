package main

import (
	"fmt"
)

type myInt int

var x myInt

func main() {
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
