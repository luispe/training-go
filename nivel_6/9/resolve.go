package main

import (
	"fmt"
)

func main() {
	f := func(msj string) string {
		return msj
	}

	msj := foo(f, "Hola LuisPe")
	fmt.Println(msj)
}

func foo(f func(msj string) string, s string) string {
	return f(s)
}
