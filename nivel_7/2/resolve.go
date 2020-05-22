package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) change(name string) {
	p.name = name
}

func main() {
	p := person{
		name: "Homer",
	}
	fmt.Println(p.name)
	fmt.Println(&p.name)

	p.change("Barnie")
	fmt.Println(p.name)
	fmt.Println(&p.name)
}
