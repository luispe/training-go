package main

import (
	"fmt"
)

func main() {
	a := "guitarra"
	if a == "piano" {
		fmt.Println("Es un piano")
	} else if a == "violin" {
		fmt.Println("Es un violin")
	} else {
		fmt.Println("Es un instrumento increíble")
	}
}
