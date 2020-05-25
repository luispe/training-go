package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("I'm a anonymous func :)")
	}()
}
