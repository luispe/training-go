package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName       string
	LastName        string
	FavoritesThings []string
}

func main() {
	p1 := person{
		FirstName:       "Homer",
		LastName:        "Simpson",
		FavoritesThings: []string{"Beer", "Donuts", "Sofa"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Printf("ERROR JSON Marshal %v", err)
		return
	}
	fmt.Println(string(bs))

}
