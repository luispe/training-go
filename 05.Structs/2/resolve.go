package main

import (
	"fmt"
)

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {
	p1 := person{
		firstName:        "Homer",
		lastName:         "Simpson",
		favoriteIceCream: []string{"chocolate", "oreo cream"},
	}

	p2 := person{
		firstName:        "Maggie",
		lastName:         "Simpson",
		favoriteIceCream: []string{"strawberry", "cinnamon"},
	}

	mapPerson := map[string]person{
		p1.firstName + p1.lastName: p1,
		p2.firstName + p2.lastName: p2,
	}

	for _, v := range mapPerson {
		fmt.Println(v.firstName, v.lastName)
		for i, v := range v.favoriteIceCream {
			fmt.Println(" ", i, v)
		}
		fmt.Println("-------")
	}
}
