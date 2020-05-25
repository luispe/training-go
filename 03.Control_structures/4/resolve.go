package main

import (
	"fmt"
)

func main() {
	age := 1984
	for {
		if age == 2020 {
			fmt.Println(age)
			break
		}
		fmt.Println(age)
		age++
	}
}
