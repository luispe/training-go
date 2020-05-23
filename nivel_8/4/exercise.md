#### Ejercicio #4

Comenzando con este código
```
package main

import (
	"fmt"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"Una", "persona", "que", "nunca", "ha", "cometido", "un", "error", "nunca", "intenta", "nada", "nuevo"}

	fmt.Println(xi)
	// sort slice int here
	fmt.Println(xi)

	fmt.Println(xs)
	// sort slice string here
	fmt.Println(xs)
}
```
Ordena el slice de int y el slice de string utilizando la standard lib de go.