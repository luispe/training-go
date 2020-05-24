#### Ejercicio #3

Comenzando con este código:
```
package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("Finish program.")
}

func gen() <-chan int {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		c <- i
	}

	return c
}
```

1. Extrae los valores del canal usando un ciclo for range