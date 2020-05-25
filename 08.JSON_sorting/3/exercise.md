#### Ejercicio #3

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
Codifica a JSON el []user enviando el resultado a Stdout.
Pista: Necesitarás usar json.NewEncoder(os.Stdout).encode(v interface{})