package main

import (
	"fmt"
)

type Human interface {
	say()
}

type Person struct {
	Name string
}

func (p *Person) say() {
	fmt.Println("Hi! i'm ", p.Name)
}

func someSay(h Human) {
	h.say()
}

func main() {
	p1 := Person{
		Name: "Homer",
	}

	//someSay(&p1)
	//someSay(p1) resolve quiz 5b, error parameter, program not compile
	p1.say()
}
