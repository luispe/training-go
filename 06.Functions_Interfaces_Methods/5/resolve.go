package main

import (
	"fmt"
	"math"
)

type Square struct {
	Base float64
}

func (s Square) Area() float64 {
	return s.Base * s.Base
}

type Circle struct {
	Radius float64
}

type Form interface {
	Area() float64
}

func info(f Form) {
	fmt.Println(f.Area())
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	c := Circle{
		Radius: 11.02,
	}

	sq := Square{
		Base: 11,
	}

	info(c)
	info(sq)
}
