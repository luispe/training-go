#### Ejercicio #1

Comenzando con este código
```
package main

type user struct {
    name string
    age   int
}

func main() {
    u1 := user{
        name: "Homer",
        age:   42,
    }

    u2 := user{
        name: "Barnie",
        age:   39,
    }

    u3 := user{
        name: "Moe",
        age:   54,
    }

    users := []user{u1, u2, u3}

    // code here
}
```
Utiliza Marshal de la standard lib de golang para transformar el slice de users a JSON. Hay una pequeña curva de
dificultad en el ejercicio. Recuerda que necesitas para exportar un valor de un paquete.