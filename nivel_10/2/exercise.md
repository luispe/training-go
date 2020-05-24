#### Ejercicio #2

Haz que funcionen los siguientes códigos:

1
```
package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
```
2
```
package main

import (
	"fmt"
)

func main() {
	cr := make(<-chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
```