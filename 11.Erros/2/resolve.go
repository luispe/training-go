package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

}

// toJSON también necesita retorna un error
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("ERROR convert to JSON %v", err)
	}
	return bs, nil
}
