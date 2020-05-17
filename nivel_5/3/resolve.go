package main

import (
	"fmt"
)

type car struct {
	doors string
	color string
}

type truck struct {
	car
	fourWheels bool
}

type sedan struct {
	car
	luxuries bool
}

func main() {
	t := truck{
		car: car{
			doors: "2",
			color: "black",
		},
		fourWheels: true,
	}
	s := sedan{
		car: car{
			doors: "4",
			color: "white",
		},
		luxuries: true,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println("Doors truck", t.doors)
	fmt.Println("Doors sedan", s.doors)
}
