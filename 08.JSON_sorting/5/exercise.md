#### Ejercicio #5

Comenzando con este código
```
package main

import (
	"fmt"
)

type user struct {
	FirstName       string
	LastName        string
	Age             int
	FavoritesThings []string
}

func main() {
	u1 := user{
		FirstName: "Homer",
		LastName:  "Simpson",
		Age:       42,
		FavoritesThings: []string{
			"Beer",
			"Donuts",
			"Sofa",
		},
	}

	u2 := user{
		FirstName: "Barnie",
		LastName:  "Gomez",
		Age:       39,
		FavoritesThings: []string{
			"Beer",
			"Helicopter",
			"Bar",
		},
	}

	u3 := user{
		FirstName: "Ned",
		LastName:  "Flanders",
		Age:       40,
		FavoritesThings: []string{
			"Church",
			"Family",
			"Walk",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// code here

}
```
1. Ordena el slice de user por:
    * Age
    * LastName
2. Ordena el slice de string "FavoritesThings" para cada user.
    * Imprime todo de una manera agradable.