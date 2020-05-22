package main

import (
	"fmt"
	"math"
)

type square struct {
	base float64
}

func (s square) area() float64 {
	return s.base * s.base
}

type circle struct {
	radius float64
}

type form interface {
	area() float64
}

func info(f form) {
	fmt.Println(f.area())
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := circle{
		radius: 11.02,
	}

	sq := square{
		base: 11,
	}

	info(c)
	info(sq)
}
