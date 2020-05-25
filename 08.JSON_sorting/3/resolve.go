package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	FirstName       string
	LastName        string
	Age             int
	FavoritesThinks []string
}

func main() {
	u1 := user{
		FirstName: "Homer",
		LastName:  "Simpson",
		Age:       42,
		FavoritesThinks: []string{
			"Beer",
			"Donuts",
			"Sofa",
		},
	}

	u2 := user{
		FirstName: "Barnie",
		LastName:  "Gomez",
		Age:       39,
		FavoritesThinks: []string{
			"Beer",
			"Helicopter",
			"Bar",
		},
	}

	u3 := user{
		FirstName: "Ned",
		LastName:  "Flanders",
		Age:       40,
		FavoritesThinks: []string{
			"Church",
			"Family",
			"Walk",
		},
	}

	users := []user{u1, u2, u3}

	if err := json.NewEncoder(os.Stdout).Encode(users); err != nil {
		fmt.Println("ERROR New Encode []user, error: ", err)
	}
}
