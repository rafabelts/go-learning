# Escribir tests
Escribir un test es como escribir una funcion pero con ciertas reglas

- En go debe ir en un archivo con nombre xxx_test.go
- La funcion test debe iniciar con la palabra Test
- La funcion solo toma el argumento <b>t *testing.T</b>
- Para usar el tipo *testing.T, se necesida importar la libreria "testing", como se hace con el fmt


## Subtests
```go
t.Run("descripcion", func(t *testing.T) {
    cuerpo del subtest
})    
```
En ocasiones es util agrupar los tests alrededor de "una cosa" y despues tener subtests describiendo diferentes escenarios

```go
func assertCorretMessage(t testing.TB, got, want string) {
    t.Helper()
    if(got != want) {
        message
    } 
}    
```
Aqui se refactoriza el mensaje de asercion en una funcion para evitar duplicar y mejorar la redaccion de los test.
Se necesita pasar el argumento t *testing.T asi podemos decir al codigo q fallo cuando se necesita

Para estas funciones helper es bueno aceptar un tipo testing.TB el cual es una interfaz que *testing.T
y *testing.B satisfacen, asi puede ser utilizada desde un test o benchmark

t.Helper() es necesitado para indicar a la suite del test que el metodo es un helper. Al hacer esto, cuando
falla, el numero de linea reportado sera en nuestro llamada a la funcion en lugar de en el helper. Esto ayuda a 
encontrar problemas de manera mas sencilla
