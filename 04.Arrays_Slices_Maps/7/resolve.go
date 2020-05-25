package main

import (
	"fmt"
)

func main() {
	k := []string{"Batman", "Jefe", "Vestido de negro"}
	l := []string{"Robin", "Ayudante", "Vestido de colores"}

	x := [][]string{k, l}

	for i, sl := range x {
		fmt.Println("Registro: ", i)
		for j, v := range sl {
			fmt.Printf("\tÍndice: %d\tValor: %s\n", j, v)
		}
	}
}
