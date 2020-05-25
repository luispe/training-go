#### Ejercicio #1

Haz que funcione el siguiente código:
```
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
```
1. Usando una función literal.
2. Usando buffer channel.