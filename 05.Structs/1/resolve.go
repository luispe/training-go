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

	fmt.Println(p1)
	for _, favIC := range p1.favoriteIceCream {
		fmt.Println(favIC)
	}
	fmt.Println(p2)
	for _, favIC := range p2.favoriteIceCream {
		fmt.Println(favIC)
	}
}
