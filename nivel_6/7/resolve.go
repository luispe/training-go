package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("I'm anonymous func :)")
	}
	f()
	fmt.Printf("%T", f)
}
