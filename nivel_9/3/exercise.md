#### Ejercicio #3

1. Usando goroutines, crea un programa incremento.
    * Haz que una variable guarde un valor incremento.
    * Lanza varias goroutines.
        + Leer el valor del incremento.
            * Almacenarlo en una nueva variable.
        + Ceder el procesador con runtime.Gosched()
        + Incrementa la nueva variable.
        + Escribe el valor en la nueva variable de vuelta a la variable incremento.
2. Usa waitGroups para esperar que todas la goroutines finalicen
3. Lo anterior generará un race condition.
4. Comprueba que es una race condition usando el flag -race.
5. Si necesitas ayuda, aquí tienes una [pista](https://play.golang.org/p/a-tdD-7lTId)