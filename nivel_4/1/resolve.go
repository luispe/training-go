package main

import (
	"fmt"
)

func main() {
	k := [5]int{1, 2, 3, 4, 5}

	for _, v := range k {
		fmt.Println(v)
	}
	fmt.Printf("%T", k)
}
