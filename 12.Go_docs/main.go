package main

import (
	"fmt"

	"github.com/luispe/training-go/12.Go_docs/dog"
)

type can struct {
	Name string
	Age  int
}

func main() {
	c := can{
		Name: "Firulais",
		Age:  dog.AgeDog(2),
	}
	fmt.Printf("Name: %s\nAge: %d", c.Name, c.Age)
}
