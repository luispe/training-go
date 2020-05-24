#### Ejercicio #4

Comenzando con este código:
```
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

func gen(salir chan<- int) <-chan int {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		c <- i
	}

	return c
}
```

1. Extrae los valores del canal usando una declaración select.