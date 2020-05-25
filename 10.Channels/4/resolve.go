package main

import (
	"fmt"
)

func main() {
	exit := make(chan int)
	c := gen(exit)

	receive(c, exit)

	fmt.Println("Finish program.")
}

func gen(exit chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		exit <- 1
		close(c)
	}()

	return c
}

func receive(c, exit <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-exit:
			return
		}
	}
}
