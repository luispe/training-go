#### Ejercicio #2

Comenzando con este código
```
package main

import (
	"fmt"
)

func main() {
	s := `[{"FirstName":"Homer","LastName":"Simpson","Age":42,"FavoritesThinks":["Donuts","Beer","Sofa"]},
{"FirstName":"Barnie","LastName":"Gomez","Age":39,"FavoritesThinks":["Beer","Helicopter","Bar"]},
{"FirstName":"Ned","LastName":"Flanders","Age":40,"FavoritesThinks":["Church","Family","Walk"]}]`
	fmt.Println(s)

}
```
Utiliza unmarshal de la standard lib de golang para transformar el JSON a una estructura de datos. Puede usar esta
[herramienta online](https://mholt.github.io/json-to-go/) para tener una orientación de como crear la estructura para un JSON.