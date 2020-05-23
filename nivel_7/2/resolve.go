package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) change(name string) {
	p.Name = name
}

func main() {
	p := Person{
		Name: "Homer",
	}
	fmt.Println(p.Name)
	fmt.Println(&p.Name)

	p.change("Barnie")
	fmt.Println(p.Name)
	fmt.Println(&p.Name)
}
