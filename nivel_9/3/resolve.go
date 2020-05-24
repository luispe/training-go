package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	gs := 50
	inc := 0
	var incGo int
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			incGo = inc

			runtime.Gosched()
			incGo++
			incGo = inc
			fmt.Println(inc)
		}()
	}
	fmt.Println("El valor final de incremento es: ", inc)
}
