package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "Hola mundo"
}
