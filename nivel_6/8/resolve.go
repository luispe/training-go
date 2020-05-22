package main

import (
	"fmt"
)

func main() {
	f := foo()
	f()
}

func foo() func() {
	return func() {
		fmt.Println("I'm foo and return a func, call me and execute")
	}
}
