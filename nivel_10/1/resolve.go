package main

import (
	"fmt"
)

func main() {
	// Resolve 2 use buffer channel
	c := make(chan int, 1)
	c <- 42

	/* Resolve 1.
	go func() {
		c <- 42
	}()
	*/

	fmt.Println(<-c)
}
