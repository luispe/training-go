package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName       string
	LastName        string
	Age             int
	FavoritesThinks []string
}

func main() {
	s := `[{"FirstName":"Homer","LastName":"Simpson","Age":42,"FavoritesThinks":["Donuts","Beer","Sofa"]},
{"FirstName":"Barnie","LastName":"Gomez","Age":39,"FavoritesThinks":["Beer","Helicopter","Bar"]},
{"FirstName":"Ned","LastName":"Flanders","Age":40,"FavoritesThinks":["Church","Family","Walk"]}]`

	var persons []Person
	if err := json.Unmarshal([]byte(s), &persons); err != nil {
		fmt.Println("Error unmarshal data", err)
		return
	}
	fmt.Println(persons)

	for i, v := range persons {
		fmt.Println("Personaje #", i+1)
		fmt.Println("\t", v.FirstName, v.LastName, v.Age)
		for _, v := range v.FavoritesThinks {
			fmt.Println("\t\t", v)
		}
	}
}
