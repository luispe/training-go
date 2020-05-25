package main

import (
	"fmt"
)

func main() {
	k := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(k[:5])
	fmt.Println(k[5:])
	fmt.Println(k[2:7])
	fmt.Println(k[1:6])
}
