# Table driven test
Son utiles cuando se quiere construir una lista de casos que pueden ser testeados de
la misma manera

```go

func TestArea(t *testing.T) {
    areaTests := []struct {
        shape Shape
        want float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
    }
    
    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

}
```

Se crea un struct anonimo, llamado area tests

Se declara un slice de structs usando []struct con dos campos, el de figura y
el de querer. Despues se llena con los casos

Es muuy facil agregar una nueva figura, implementar el area y agregarla a los casos para
probar. De igual forma, si un bug se encuentra en Area sera mas facil agregar un nuevo
test para probarlo antes de arreglarlo

Este estilo de test es muy bueno cuando queremos probar varias implementaciones de una
interfaz, o si los datos pasados en una funcion tiene varios requerimientos que se deben de probar

