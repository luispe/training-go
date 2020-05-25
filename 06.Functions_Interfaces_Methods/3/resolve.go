package main

import (
	"fmt"
)

type person struct {
	firstName, lastName string
	Age                 int
}

func (p person) hi() {
	fmt.Printf("Hi! my name is %s & i'm %d years old.", p.firstName, p.Age)
}

func main() {
	p := person{
		firstName: "LuisPe",
		lastName:  "Toloy",
		Age:       36,
	}
	p.hi()
}
