package main

import (
	"fmt"
)

func main() {
	deporteFav := "futbol"
	switch deporteFav {
	case "futbol":
		fmt.Println("Diego Armando Maradona")
	case "basquet":
		fmt.Println("Michael Jordan")
	case "automovilismo":
		fmt.Println("Juan Manuel Fangio")
	}
}
