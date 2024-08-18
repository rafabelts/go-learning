/*
En un programa de go se tendra un paquete main definido con
una func main dentro.

Los paquetes son una forma de agrupar codigo de Go relacionado
*/

package main

// import "fmt" // Importa el paquete Println para mostrar a consola


// Se pueden agrupar las constantes en un solo bloque
const (
    spanish = "Spanish"
    french = "French"

    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix =  "Bonjour, " 
)

func Hello(name string, language string) string {
    helloPrefix := getPrefix(language)
 
    if (name == "" || name == " ") {
        name = "World"
    }

    return helloPrefix + name
}

/*
- Se utiliza un valor de retorno con nombre: (prefix string)

- Esto crea una variable llamada prefix en la funcion:
    - A este le asigna un valor 0
        - Se puede devolver lo que esta puesto llamando return en vez de tener que hacer un return prefix

- La funcion inicia con minusculas. En Go, las funciones publicas inician con mayusculas, y las privadas con
minuscula. No queremos que funciones internas del algoritmo queden expuestas al mundo, por eso se hace privada la funcion
*/

func getPrefix(language string) (prefix string) {
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    return
}
