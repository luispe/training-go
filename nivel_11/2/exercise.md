#### Ejercicio #2

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

	bs, err := toJSON(p1)

	fmt.Println(string(bs))

}

// toJSON también necesita retorna un error
func toJSON(a interface{}) []byte {
	bs, err := json.Marshal(a)
}
```
Crea un mensaje de error personalizado usando “fmt.Errorf”.