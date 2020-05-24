package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var inc int64
	gs := 50
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&inc, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&inc))
		}()
	}

	wg.Wait()
	fmt.Println("El valor final de incremento es: ", inc)
}
