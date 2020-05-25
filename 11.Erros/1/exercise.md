#### Ejercicio #1

Comienza con este código:
```
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName       string
	LastName        string
	FavoritesThinks []string
}

func main() {
	p1 := person{
		FirstName:       "Homer",
		LastName:        "Simpson",
		FavoritesThinks: []string{"Beer", "Donuts", "Sofa"},
	}

	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

}
```
En vez de usar el identificador blank (underscore), asegúrate de que el código esté chequeando y manejando el error.