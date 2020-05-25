package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 11
	}()

	if v, ok := <-c; ok {
		fmt.Println("Access read channel", v, ok)
	}

	close(c)

	if v, ok := <-c; !ok {
		fmt.Println("NOT access read channel, chanell close", v, ok)
	}
}
