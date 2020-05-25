package main

import (
	"fmt"
)

const gs = 10

func main() {
	c := make(chan int)

	for i := 0; i < gs; i++ {
		go func() {
			for i := 0; i < gs; i++ {
				c <- i
			}
		}()
	}
	for i := 0; i < gs; i++ {
		fmt.Println(i, <-c)
	}
}
