package main

import (
	"fmt"
)

func main() {
	prov := []string{"Buenos Aires", "Catamarca", "Chaco", "Chubut", "Córdoba", "Corrientes", "Entre Ríos", "Formosa", "Jujuy", "La Pampa", "La Rioja", "Mendoza", "Misiones", "Neuquén", "Río Negro", "San Juan", "San Luis", "Santa Cruz", "Santa Fe", "Santiago del Estero", "Tierra del Fuego", "Tucuman"}
	fmt.Printf("La longitud del slice es de: %d\n", len(prov))
	fmt.Printf("La capacidad del slice es de: %d\n", cap(prov))

	fmt.Println("Las provincias de Argentina son: ")
	for i := 0; i < len(prov); i++ {
		fmt.Println(prov[i])
	}
}
