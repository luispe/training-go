#### Ejercicio #2

1. Crea un struct persona
2. Crea una función llamada "change" con una *persona como parámetro
    * Cambia el valor almacenado en la dirección de memoria del tipo *persona.
    * El compilador asigna valores a las variables. ¿Cómo son llamados esos valores?.
    * #### Importante
        + Para desreferenciar un struct, usa (*valor).campo
            * p1.nombre
            * (*p1).nombre
        + "Como una excepción, si el tipo de X es un tipo puntero con nombre y (*x).c es una expresión válida de
        selección denotando un campo (pero no un método), x.c es una forma abreviada de (*x).c".
3. Crea un valor de tipo persona.
    * Imprime el valor.
4. En func main
    * Llama change
5. En func main
    * Imprime el valor.