package main

import (
	"fmt"
)

func main() {
	favoritesObject := map[string][]string{
		"cosme_fulanito": {"celular", "bicicleta", "cafe"},
		"homer_simpson":  {"rosquillas", "sofa", "cerveza"},
		"bart_simpson":   {"skate", "aerosol", "TV portatil"},
	}

	favoritesObject["marge_simpson"] = []string{"collar de perlas", "auto", "vestido"}
	delete(favoritesObject, "cosme_fulanito")

	for i, p := range favoritesObject {
		fmt.Println("Registro: ", i)
		for j, favObj := range p {
			fmt.Println("\t", j, favObj)
		}
	}
}
