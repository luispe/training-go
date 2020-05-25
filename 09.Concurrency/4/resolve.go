package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	inc := 0
	gs := 50
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			incGo := inc

			incGo++
			inc = incGo
			fmt.Println(inc)
		}()
	}

	wg.Wait()
	fmt.Println("El valor final de incremento es: ", inc)
}
