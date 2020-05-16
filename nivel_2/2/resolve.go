package main

import (
	"fmt"
)

func main() {
	a := (11 == 11)
	b := (11 <= 10)
	c := (11 >= 10)
	d := (11 != 10)
	e := (11 < 10)
	f := (11 > 10)

	fmt.Println(a, b, c, d, e, f)
}
