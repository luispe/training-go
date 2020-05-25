package main

import (
	"fmt"
)

const channels = 10

func main() {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; i < channels; i++ {
			c <- i
		}
	}()

	for v := range c {
		fmt.Println(v)
	}
}
